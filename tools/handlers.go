package main

/*
Author: Jeff Berkowitz
Copyright (C) 2023 Jeff Berkowitz

This file is part of mkmd.

mkmd is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation, either version 3
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see http://www.gnu.org/licenses/.
*/

import (
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// An optable for dumping the entire document for debugging purposes
var printPass = opTable{
	defaultAction: printNode,
	typeFuncs: [6]opFunc{},
	elementPreFuncs: nil,
	elementPostFuncs: nil }

// An optable for a pass that gets .rdf files found in specific <link> tags
var rdfPass = opTable {
	defaultAction: nil,
	typeFuncs: [6]opFunc{},
	elementPreFuncs: map[
		atom.Atom]opFunc{atom.Link: rdfHandler},
	elementPostFuncs: nil,
}

// An optable for generating .md files from the Wiki's HTML files
var mdPass = opTable {
	defaultAction: nil,
	typeFuncs: [6]opFunc{nil, doText, nil, nil, doComment, doDocType},
	elementPreFuncs: map[atom.Atom]opFunc{
		atom.A:  doAtagOpen,
		atom.H1: doHeaderOpen,
		atom.H2: doHeaderOpen,
		atom.H3: doHeaderOpen,
		atom.H4: doHeaderOpen,
		atom.H5: doHeaderOpen,
		atom.H6: doHeaderOpen,
		atom.Img: doImgOpen,
		atom.Li: doLiOpen,
		atom.Ol: doOlOpen,
		atom.Script: doScriptOpen,
		atom.Title: doTitleOpen,
		atom.Ul: doUlOpen,
	},
	elementPostFuncs: map[atom.Atom] opFunc{
		atom.A:  doAtagClose,
		atom.Body: doHtmlClose,
		atom.H1: doHeaderClose,
		atom.H2: doHeaderClose,
		atom.H3: doHeaderClose,
		atom.H4: doHeaderClose,
		atom.H5: doHeaderClose,
		atom.H6: doHeaderClose,
		atom.Html: doHtmlClose,
		atom.Li: doLiClose,
		atom.Ol: doOlClose,
		atom.Script: doScriptClose,
		atom.Title: doTitleClose,
		atom.Ul: doUlClose,
	},
}

// =============================================================
// Operation functions
// =============================================================

func prototype(n *html.Node, cx *context) error {
	return nil // copy this to make new handlers
}

/*

This is an example of a reoccurring block that I need to learn how to treat.
It's got three divs with recognizable classes containing an A tag that links
to a description page (an HTML page despite the name that ends with .png).

Then there's an img tag in the A tag with a link to a legitimate image.

<div class="center">
    <div class="thumb tnone">
        <div class="thumbinner" style="width:280px;">
            <a href="/wiki/index.php?title=File:NES-2A03-decimal-DAA-removed.png" class="image">
                <img alt="" src="/wiki/images/8/89/NES-2A03-decimal-DAA-removed.png" width="278" height="200" class="thumbimage" />
            </a>
            <div class="thumbcaption">
                Transistor t2556 in NES 2A03
            </div>
        </div>
    </div>
</div>

There's a similar bunch of code distinguished by a link to the currently
non-existent (in the wiki target directory) url skins/common/images/magnify-clip.png
that also needs to be addressed.

*/
// According to the standard, A-tags don't nest
func doAtagOpen(n *html.Node, cx *context) error {
	href := getAttrVal(n, "href")
	if len(href) == 0 {
		fmt.Fprintln(os.Stderr, "A tag with no href")
		return nil
	}

	cx.emitString("\n[")
	cx.atagRemainder = "](" + urlSafeUrl(href) + ")"
	return nil
}

func doAtagClose(n *html.Node, cx *context) error {
	cx.emitString(cx.atagRemainder)
	return nil
}

// </html> endtag. Write the output file.
func doHtmlClose(n *html.Node, cx *context) error {
	outDir := cx.outputDirectory
	if len(outDir) == 0 {
		outDir = "."
	}
	inFileName := path.Base(cx.inputFilePath)
	outPath := path.Join(outDir, inFileName) + ".md"
	err := os.WriteFile(outPath, []byte(cx.md.String()), 0644)
	if err != nil {
		fatal("write result to \"%s\" failed: %v", outPath, err)
	}
	return nil
}

func doImgOpen(n *html.Node, cx *context) error {
	imgText := getAttrVal(n, "alt")
	if len(imgText) == 0 {
		imgText = "Image (no description given)"
	}

	imgLink := getAttrVal(n, "src")
	if len(imgLink) == 0 {
		fmt.Fprintln(os.Stderr, "Warning: img tag with no src")
		return nil
	}

	// "If s doesn't start with prefix, s is returned unchanged."
	imgLink = strings.TrimPrefix(imgLink, "/wiki/")
	cx.emitString("\n![%s](%s)\n", imgText, imgLink)
	return nil
}

func doHeaderOpen(n *html.Node, cx *context) error {
	const hashes = "#######"
	cx.emitString("\n" + hashes[0:1 + int(n.Data[1] - '0')] + " ")
	return nil
}

func doHeaderClose(n *html.Node, cx *context) error {
	cx.emitString("\n")
	return nil
}

func doLiOpen(n *html.Node, cx *context) error {
	cx.emitString("\n%s ", cx.liString)
	return nil
}

func doLiClose(n *html.Node, cx *context) error {
	return nil
}

func doTitleOpen(n *html.Node, cx *context) error {
	cx.emitString("\n# ")
	return nil
}

func doOlOpen(n *html.Node, cx *context) error {
	cx.liString = "1."
	return nil
}

func doOlClose(n *html.Node, cx *context) error {
	cx.liString = ""
	return nil
}

func doTitleClose(n *html.Node, cx *context) error {
	cx.emitString("\n")
	return nil
}

func doDocType(n *html.Node, cx *context) error {
	cx.emitString("**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**\n")
	return nil
}

func doText(n *html.Node, cx *context) error {
	s := strings.TrimSpace(n.Data)
	if len(s) != 0 {
		cx.emitString(n.Data)
	}
	return nil
}

func doUlOpen(n *html.Node, cx *context) error {
	cx.liString = "-"
	return nil
}

func doUlClose(n *html.Node, cx *context) error {
	cx.liString = ""
	return nil
}

func doComment(n *html.Node, cx *context) error {
	// The Wayback Machine emits this content before its footer.
	// The WM Downloader is documented as only downloading the
	// page "as it was", but apparently this does not work because
	// the downloaded pages have all this WM footer material.
	// There is no matching enableOutput() for this - it continues
	// to the end of the file, when we emit our footer instead.
	if strings.HasPrefix(n.Data, " Saved in parser cache") {
		cx.disableOutput() // all the way to the end
	}
	return nil
}

func doScriptOpen(n *html.Node, cx *context) error {
	cx.disableOutput()
	return nil
}

func doScriptClose(n *html.Node, cx *context) error {
	cx.enableOutput()
	return nil
}

// End of support for -m. Begin support for -r.

// rdfHandler is invoked for <link> tags only. It identifies tags
// that link .rdf files containing authorship and license data and
// downloads their latest version.
func rdfHandler(n *html.Node, cx *context) error {
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

	// If we have the correct kind of link, fetch it. Note: We want to
	// maintain certain correspondences between file names. But the Jekyll
	// processor (or some part of the Gitlab Pages pipeline) is not happy
	// with filenames having illegal URL characters like single quotes or
	// parens, and three of the top level files (and their corresponding
	// .rdf files) contain single quotes. Yet all we have here is the href
	// to the file, which does not contain a single quote - it contains an
	// URL-encoded single quote, %27. Eventually I made the decision to
	// create a substitution rule and rename all the files. This means I
	// need to make the same substitutions in the URLs that refer to them.
	// So I need to get the filename component of the URL with the %27
	// expanded to a single quote, then run the substitution rule on the
	// the resulting string, then write the file under that name, and then
	// finally (later, not here) fix the href to match. The substitution
	// rule for characters produces only URL-safe characters, so there is
	// never any need to URL encode after running the rule.
	if isRDF && len(href) > 0 {
		url, err := getMostRecentUrl(makeWaybackApiQuery(href))
		if err != nil {
			return err
		}
		url, err = fixupForRdf(url)
		if err != nil {
			return err
		}
		page, err := getBody(url)
		if err != nil {
			return err
		}
		rdfName, err := getTitle(href)
		if err != nil {
			return err
		}
		rdfName = urlSafeName(rdfName)
		makeOutputDir(cx.outputDirectory)
		if err = os.WriteFile(path.Join(cx.outputDirectory, rdfName + ".rdf"), page, 0600); err != nil {
			return err
		}
	}

	return nil
}

// End of support for -r. Begin general handlers.

// An opFunc that returns an internal error
func internalError(n *html.Node, cx *context) error {
	return fmt.Errorf("internal error: type %v not expected (context %v)", n, cx)
}

// An opFunc that prints "not handled: thing" for use as a default
func notHandled(n *html.Node, cx *context) error {
	dbg("not handled: node %v (context %v)\n", n, cx)
	return nil
}

// A debugging opFunc that just prints the node with indent
func printNode(n *html.Node, cx *context) error {
	fmt.Printf("%*sType=%s DataAtom=%v Data=%v Attr=%v\n", cx.depth*2, "",
		typeNames[n.Type], n.DataAtom, strings.TrimSpace(n.Data), n.Attr)
	return nil
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
