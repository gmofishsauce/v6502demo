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
		to download it. The downloaded XML (RDF) file is given a safe name.
	-m
		Write a markdown (".md") file corresponding to the document, which
		must be an HTML file. The file is written in the directory given by
		the -o argument, if suppied, default "."
	-u
		Update the filename of the argument file to a safe name.
	
	-x
		This flag is a hack to enable temporary types of debugging.

Files downloaded by the Wayback Machine Downloader from the visual6502.org
MediaWiki in the Wayback Machine have filenames derived from the titles of
the MediaWiki pages. These titles may contain any character: they are not URL
safe. The Gitlab Pages pipeline uses the Jekyll tool to create a static site,
and the pipeline does not tolerate URL-unsafe characters. Therefore, it's
necessary to rename the downloaded files. The renaming process must not break
links between pages, and the links are (as far as I know) URL encoded. So a
file with a single quote in the name will have a %27 in an URL that links to
it. Since the single quote in the filename must be changed to an URL safe
character, it becomes unnecessary (but I hope harmless) to URL encode links 
to the file. The solution I use is to replace URL-unsafe characters in file
names with URL safe characters (dash and tilde). When reprocessing URLs, the
URL is decoded to regenerate the URL-unsafe characters and the URL (really
just the name component) is subjected to the same replacement rule.

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
	xflag := flag.Bool("x", false, "dump information about argument file")
	flag.Parse()

	// Now make a table that maps the value of an operation flag,
	// true or false, to the operation table that implements it.
	flagOpTable := []flagOp{
		{*dflag, &printPass, "-d"},
		{*rflag, &rdfPass, "-r"},
		{*mflag, &mdPass, "-m"},
		{*xflag, &xPass, "-x"},
	}

	file := flag.Args()
	if len(file) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: mkmd flags file")
		flag.PrintDefaults();
		os.Exit(1)
	}

	if *uflag {
		if err := renameFileToUrlSafe(file[0]); err != nil {
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

	dbg("Parsing %s", name)
	doc, err := html.Parse(f)
	if err != nil {
		fatal(err.Error())
	}

	processed := 0

	for _, flop := range flagOpTable {
		if flop.flagVal {
			dbg("Processing %s", flop.name)
			setOpTable(flop.ops)
			cx := NewContext(*oflag, name)
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
		if err != nil {
			return err
		}
		cx.depth--
	}
	if post := nodeToOp(n, false); post != nil {
		if err = post(n, cx); err != nil {
			return err
		}
	}
	return nil
}

