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
	},
	elementPostFuncs: map[atom.Atom] opFunc{
	},
}

func proto(n *html.Node, cx *context) error {
	return nil
}

func doDocType(n *html.Node, cx *context) error {
	return nil
}

func doText(n *html.Node, cx *context) error {
	return nil
}

// =============================================================
// Operation functions
// =============================================================

// An opFunc that returns an internal error
func internalError(n *html.Node, cx *context) error {
	return fmt.Errorf("internal error: type %v not expected (context %v)", n, cx)
}

// An opFunc that prints "not handled: thing" for use as a default
func notHandled(n *html.Node, cx *context) error {
	fmt.Printf("not handled: node %v (context %v)\n", n, cx)
	return nil
}

// A debugging opFunc that just prints the node with indent
func printNode(n *html.Node, cx *context) error {
	fmt.Printf("%*sType=%s DataAtom=%v Data=%v Attr=%v\n", cx.depth*2, "",
		typeNames[n.Type], n.DataAtom, strings.TrimSpace(n.Data), n.Attr)
	return nil
}

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

