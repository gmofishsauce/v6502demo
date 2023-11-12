/*
mkmd is a catch-all tool that performs various tasks on documents
downloaded from a website, specifically a MediaWiki site. Mkmd
is usually run from a script or find(1) invocation; it processes
a single document on each execution.

Usage:

    mkmd [flags] document-path

The flags are:

	-d
		Dump the HTML parse tree of the HTML document to stdout
	-o directory
		If any downloads occur or markdown files are generated, store the
		downloaded or generated file(s) in `directory`, which is created
		if it does not exist. If no downloads or file generation occur,
		this option does nothing. The default output directory is "."
    -r
		Scan the document for a `<link>` tag with a `type` attribute having
		value `application/rdf+xml`. If found, use the Wayback Machine API to
		construct a URL that is likely to reference the RDF file and attempt
		to download it.
	-m
		Write a markdown (".md") file corresponding to the document, which must
		be an HTML file. The file is written in the directory given by -o.
	-u
		Update the filename of the argument file to contain no shell metacharacters
		or reserved URL characters. Details are TBD.


*/
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
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type flagOp struct {
	flagVal bool
	ops *opTable
	name string
}

func main() {
	dflag := flag.Bool("d", false, "dump html")
	rflag := flag.Bool("r", false, "get rdf content")
	oflag := flag.String("o", ".", "set output directory")
	mflag := flag.Bool("m", false, "create markdown file")
	uflag := flag.Bool("u", false, "replace non-URL characters in file name")
	flag.Parse()

	// Now make a table that maps the value of an operation flag,
	// true or false, to the operation table that implements it.
	flagOpTable := []flagOp{
		{*dflag, &printPass, "-d"},
		{*rflag, &rdfPass, "-r"},
		{*mflag, &mdPass, "-m"},
	}

	file := flag.Args()
	if len(file) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkmd flags file")
		flag.PrintDefaults();
		os.Exit(1)
	}

	if *uflag {
		if err := makeUrlSafeFileName(file[0]); err != nil {
			fatal(err.Error())
		}
		fmt.Fprintf(os.Stderr, "mkmd: success\n")
		os.Exit(0)
	}

	name := file[0]
	f, err := os.Open(name)
	if err != nil {
		fatal(err.Error())
	}
	defer f.Close()

	dbg("Parsing %s\n", name)
	doc, err := html.Parse(f)
	if err != nil {
		fatal(err.Error())
	}

	processed := 0

	for _, flop := range flagOpTable {
		if flop.flagVal {
			dbg("Processing %s\n", flop.name)
			setOpTable(flop.ops)
			cx := &context{0, *oflag, name}
			if err = process(doc, cx); err != nil {
				fatal(err.Error())
			}
			processed++
		}
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

