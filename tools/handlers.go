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
	typeFuncs: [6]opFunc{nil, doText, nil, nil, nil, doDocType},
	elementPreFuncs: map[atom.Atom]opFunc{
		atom.Img: doImg,
	},
	elementPostFuncs: map[atom.Atom] opFunc{
	},
}

// An optable for fixing filenames having inconvenient characters
var urlPass = opTable {
	defaultAction: nil,
	typeFuncs: [6]opFunc{nil, nil, nil, nil, nil, urlDocType},
	elementPreFuncs: map[atom.Atom]opFunc{
		atom.A: urlA,
		atom.Link: urlLink,
		atom.Img: urlImg,
	},
	elementPostFuncs: map[atom.Atom] opFunc{
	},
}

// =============================================================
// Operation functions
// =============================================================

func prototype(n *html.Node, cx *context) error {
	return nil
}

// The urlTag functions implement -u. This pass is intended to clean up
// shell metacharacters and non-URL characters that appear in the names
// of Wiki pages. These are a problem because the Jekyll markdown to HTML
// processor cannot successfully process these. (To be clear, I don't fully
// understand the issue - rather I just observe that removing files with
// these characters causes Github Pages builds to succeed. I don't know
// exactly who or what is at fault.)
//
// These functions don't work like other functions in this tool: instead of
// actually taking an action, they emit shell commands that will take the
// action when invoked. There are only a few of these files (for example,
// three top-level HTML files have single quotes in their file name). It's
// much safer to emit the commands for perusal than to do the work. Most of
// the work is fixing URLs that point to files with names that have changed.
//
// To fix all the relevant URLs, the code needs to find the root of the entire
// wiki. It does this by seeking a file with a specific name that should be
// present in the root of the wiki: "index.php?title=Special:AllPages" and
// using the directory holding that file as the wiki root.

const pageExpectedInRoot = "index.php?title=Special:AllPages"
var rootDirName string // XXX should go in context

func getWikiRoot(filePath string, rootIndicator string) string {
	return "TODO"
}

func urlDocType(n *html.Node, cx *context) error {
	dbg("Processing rename for %s\n", cx.fileName)
	rootDirName := getWikiRoot(cx.fileName, pageExpectedInRoot)
	if len(rootDirName) == 0 {
		fatal("Cannot find root of wiki: no directory in path \"%s\" contains \"%s\"\n",
			cx.fileName, pageExpectedInRoot)
	}
	return nil
}

func urlA(n *html.Node, cx *context) error {
	return nil
}

func urlImg(n *html.Node, cx *context) error {
	return nil
}

func urlLink(n *html.Node, cx *context) error {
	return nil
}

/*

This was how I decided that the only HTML tags with attributes
containing URLs I had to worry about were <a>, <img>, and <link>.
The code is no longer needed but kept for now in case I want to
run something like it again.

var nodesHavingURLs = make(map [string]bool, 10)

func findURLs(n *html.Node, cx *context) error {
    for _, a := range n.Attr {
		if strings.Contains(a.Val, "wiki") {
			kind := fmt.Sprintf("%s-%d-%s", typeNames[n.Type], n.DataAtom, n.Data)
			nodesHavingURLs[kind] = true
		}
    }
	if n.Type == html.DocumentNode {
		for k, v := range nodesHavingURLs {
			fmt.Printf("%s, %v\n", k, v)
		}
	}
	return nil
}

*/

// End of support for -u. Begin support for -m.

func doImg(n *html.Node, cx *context) error {
	imgLink := getAttrVal(n, "src")
	if len(imgLink) > 0 {
		emitString("![an image](%s)\n", imgLink)
	}
	return nil
}

func doDocType(n *html.Node, cx *context) error {
	return nil
}

func doText(n *html.Node, cx *context) error {
	return nil
}

// End of support for -m. Begin support for -r.

// rdfHandler is invoked for <link> tags only. It identifies tags
// that link .rdf files containing authorship and license data and
// downloads their latest version.
func rdfHandler(n *html.Node, cx *context) error {
	dbg("rdfHandler: attrs=%v\n", n.Attr)
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
