package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

type context struct {
	depth int
}

// linkHandler is invoked for <link> tags only. It identifies tags
// that link .rdf files containing authorship and license data and
// downloads their latest version.
func rdfGetter(n *html.Node, cx *context) error {
	dbg("linkHandler: attrs=%v\n", n.Attr)
	var isRDF bool
	var href string
	for _, a := range n.Attr {
		if a.Key == "type" && a.Val == "application/rdf+xml" {
			isRDF = true
		}
		if a.Key == "href" {
			href = a.Val
		}
	}
	// If we have the correct kind of link, fetch it.
	if isRDF && len(href) > 0 {
		url, err := getMostRecentUrl(makeWaybackApiQuery(href))
		if err != nil {
			return err
		}
		url = url + "&action=creativecommons"
		page, err := getBody(url)
		if err != nil {
			return err
		}
		if err = os.WriteFile(getTitle(href) + ".rdf", page, 0600); err != nil {
			return err
		}
	}

	return nil
}

// makeWaybackApiQuery(href string) makes a Wayback Machine query URL for the argument href
//
// TODO: this comment is specific to the first use of this function, which was for
// links to .rdf files, but the function itself is not specific to this file type.
//
// The argument (href from the link tag) looks something like this:
// /wiki/index.php?title=6502DecimalMode&action=creativecommons
// This actually references a .rdf document with authorship information
//
// The Machine's API page for finding the most recent copy is:
// https://archive.org/wayback/available?url=
//
// An example of an URL that can be successfully looked up is:
// http://www.visual6502.org/wiki/index.php?title=6502DecimalMode&action=creativecommons
// ^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// (this part is hardwired)  (url from href)
//
// ...although this worked, I think it should be URL encoded; also,
// the "action=creativecommons" is critical - that's what produces the RDF file.
//
// Sticking the two pieces together and visiting the API gives following result JSON:
//
//  {
//     "archived_snapshots" : {
//        "closest" : {
//           "available" : true,
//           "status" : "200",
//           "timestamp" : "20210405071434",
//           "url" : "http://web.archive.org/web/20210405071434/http://visual6502.org/wiki/index.php?title=6502DecimalMode"
//        }
//     },
//     "url" : "http://www.visual6502.org/wiki/index.php?title=6502DecimalMode"
//  }
//
// Note that the url returned in the result JSON **DOES NOT** produce the RDF page.
// It is necessary to MANUALLY re-append the `action=creativecommons` query string
// to produce it:
//
// wget -o wget.log -O rdf "http://web.archive.org/web/20210405071434/http://visual6502.org/wiki/index.php?title=6502DecimalMode&action=creativecommons"
//
// where the "action=" is manually attached ... finally produces the correct result in the local file "rdf".

const waybackAPI = "https://archive.org/wayback/available?url="
const wikiRoot =  "http://visual6502.org"

func makeWaybackApiQuery(href string) string {
	searchFor, err := url.JoinPath(wikiRoot, href)
	if err != nil {
		fatal("makeRdfLink: %v", err)
	}
	// We cannot use url.JoinPath here. It will join together the two URL fragments and
	// place "?url=" at the end, because it doesn't know searchFor is supposed to be a
	// query string value. I'm sure there's a right way to do this with the URL structure,
	// but it's not documented, so ...
	return waybackAPI + searchFor
}

/*
{
  "url": "http://visual6502.org/wiki/index.php?title=6502DecimalMode",
  "archived_snapshots":
  {
    "closest":
	{
	  "status": "200", "available": true,
	  "url": "http://web.archive.org/web/20210405071434/http://visual6502.org/wiki/index.php?title=6502DecimalMode",
	  "timestamp": "20210405071434"
	}
  }
}
*/

// Get the response body for the argument url. Return as a byte slice.
func getBody(url string) ([]byte, error) {
	dbg("getBody(%s)\n", url)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("getBody(): http.Get(%s): %v", url, err)
    }
	defer resp.Body.Close()
    dbg("getBody() resp: %v\n", resp)

	b, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("getBody(): httpResponse.Read(%s): %v", url, err)
    }
	return b, nil
}

// Get the Wayback Machine URL of the latest snapshot corresponding to
// the href given by the link argument. Uses the WM's "available" API.
func getMostRecentUrl(url string) (string, error) {
	b, err := getBody(url)
	if err != nil {
		return "", err
	}

	var data map[string]any
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	dbg("unmarshaled response: %v\n", data)

	archived_snapshots, ok := data["archived_snapshots"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("archived_snapshots not found in json response");
	}
	closest, ok := archived_snapshots["closest"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("closest not found in json response");
	}
	result, ok := closest["url"].(string)
	if !ok {
		return "", fmt.Errorf("url not found in json response");
	}
	return result, nil
}

func getTitle(url string) string {
	TODO()
	return "TODO"
}

func fatal(format string, args... any) {
	msg := fmt.Sprintf(format, args)
	fmt.Fprintf(os.Stderr, "mkmd: " + msg)
	os.Exit(1)
}

func main() {
	dflag := flag.Bool("d", false, "dump html")
	rflag := flag.Bool("r", false, "get rdf content")
	flag.Parse()

	files := flag.Args()

	in := os.Stdin
	name := "standard input"
	if len(files) > 0 {
		name = files[0]
		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkmd: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	}

	fmt.Printf("Parsing %s\n", name)
	doc, err := html.Parse(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mkmd: parse: %v\n", err)
		os.Exit(2)
	}

	processed := 0

	if *dflag {
		setOpTable(&printPass)
		if err = process(doc, &context{0}); err != nil {
			fmt.Fprintf(os.Stderr, "mkmd: process: %v\n", err)
			os.Exit(2)
		}
		processed++
	}

	if *rflag {
		setOpTable(&rdfPass)
		if err = process(doc, &context{0}); err != nil {
			fmt.Fprintf(os.Stderr, "mkmd: process: %v\n", err)
			os.Exit(2)
		}
		processed++
	}

	if processed > 0 {
		fmt.Fprintf(os.Stderr, "mkmd: success\n")
	} else {
		fmt.Fprintf(os.Stderr, "mkmd: no action requested\n")
	}
	os.Exit(0)
}

func process(n *html.Node, cx *context) error {
	var err error = nil
	if pre := nodeToOp(n, true); pre != nil {
		if err := pre(n, cx); err != nil {
			return err
		}
	}
	for c := n.FirstChild; err == nil && c != nil; c = c.NextSibling {
		cx.depth++
		err = process(c, cx)
		cx.depth--
	}
	if err != nil {
		return err
	}
	if post := nodeToOp(n, false); post != nil {
		if err = post(n, cx); err != nil {
			return err
		}
	}
	return nil
}

//func startEmit(np *node) {
//	switch np.name {
//	case "CharData":
//		// emitString(np.text)		
//		emitString("text")
//		np.processed = true
//	case "a":
//		emitString("[")
//	case "body":
//		np.processed = true
//	case "div":
//		// For now we don't do anything with divs.
//		np.processed = true
//	case "h1", "h2", "h3", "h4", "h5", "h6":
//		nHashes := np.name[1] - '0'
//		emitString(strings.Repeat("#", int(1 + nHashes)))
//		np.processed = true
//	case "head":
//		np.processed = true
//	case "html":
//		np.processed = true
//	case "img":
//		emitString("![an image](%s)", getAttrValue(np, "src"))
//		//emitString("image here")
//		np.processed = true
//	case "p":
//		emitString("\n")
//		np.processed = true
//	case "span":
//		np.processed = true
//	case "td":
//		emitString("|")
//		np.processed = true
//	case "title":
//		emitString("\n# ")
//		np.processed = true
//	case "tr":
//		emitString("\n\n")
//	}
//}
//
//func endEmit(np *node) {
//	switch np.name {
//	case "a":
//		emitString("](%s)", "href here") // getAttrValue(np, "href"))
//		np.processed = true
//	case "table":
//		emitString(strings.Repeat("|:---:", len(np.children[0].children)))
//		emitString("|\n")
//		np.processed = true
//	case "tr":
//		emitString("|\n")
//		np.processed = true
//	}
//}
//
//func emitString(format string, args ...any) {
//	fmt.Printf(format, args...)
//}
