/*
mkmd is a catch-all tool that performs various tasks on documents
downloaded from a website, specifically a MediaWiki site. Mkmd
is usually run from a script or find(1) invocation; it processes
a single document on each execution.

Usage:

    mkmd [flags] path

The flags are:

	-d
		Dump the HTML parse tree of HTML documents to stdout
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
		Write a markdown (".md") file corresponding to the path, which must
		be an HTML file. The file is written in the directory given by -o.

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

func main() {
	dflag := flag.Bool("d", false, "dump html")
	rflag := flag.Bool("r", false, "get rdf content")
	oflag := flag.String("o", ".", "set output directory")
	mflag := flag.Bool("m", false, "create markdown file")
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
		if err = process(doc, &context{0, *oflag}); err != nil {
			fmt.Fprintf(os.Stderr, "mkmd: process: %v\n", err)
			os.Exit(2)
		}
		processed++
	}

	if *rflag {
		setOpTable(&rdfPass)
		if err = process(doc, &context{0, *oflag}); err != nil {
			fmt.Fprintf(os.Stderr, "mkmd: process: %v\n", err)
			os.Exit(2)
		}
		processed++
	}

	if *mflag {
		setOpTable(&mdPass)
		if err = process(doc, &context{0, *oflag}); err != nil {
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

