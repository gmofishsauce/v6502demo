package main

import (
	"fmt"
	"net/url"
	"regexp"
)

// Functions that are specific to the visual6502.org website

// This comment contains useful information, but no longer corresponds to the
// partitioning of the functionality in the program. The comment is about getting
// RDF files from the href in the `<link>` tag.
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
// It is necessary to manually edit the link two ways. First the `action=creativecommons`
// query string must be added, and then the magical `if_` must be added to the 14-character
// datetime substring. This will produce the RDF:
//
// wget -o wget.log -O rdf "http://web.archive.org/web/20210405071434if_/http://visual6502.org/wiki/index.php?title=6502DecimalMode&action=creativecommons"

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
const wikiRoot =  "http://visual6502.org"

func makeAbsolute(wikiRef string) string {
	result, err := url.JoinPath(wikiRoot, wikiRef)
	if err != nil {
		fatal("could not create absolute path from site-relative href: %s %s\n", wikiRoot, wikiRef)
	}
	return result
}

const ccQS = "&action=creativecommons"
const magicIf = "if_"

// Fix up the URL of a top-level page so it references the page's licensing RDF
func fixupForRdf(url string) (string, error) {
	re := regexp.MustCompile(`.*[\d]{14}`)
	prefix := re.FindString(url)
	suffix := url[len(prefix):]
	if len(prefix) == 0 || len(suffix) == 0 {
		return "", fmt.Errorf("fixupForRdf(): failed to match URL")
	}
	return prefix+magicIf+suffix+ccQS, nil
}

// Useful information about making markdown from HTML which is
// the ultimate goal of this tooling. The functions themselves
// are completely obsolete. Each case in the switches corresponds
// to one pre- or post-element handler function.

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
