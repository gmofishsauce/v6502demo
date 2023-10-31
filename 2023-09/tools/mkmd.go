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
		If any downloads occur, then store the downloaded files in
		`directory`, which is created if it does not exist. If no
		downloads occur, this option does nothing.
    -r
		Scan the document for a `<link>` tag with a `type` attribute
		having value `application/rdf+xml`. If found, use the Wayback
		Machine API to construct a URL that is likely to reference the
		RDF file and attempt to download it.

*/
package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)


// Context passed into the tree walker functions
type context struct {
	depth int
	outputDirectory string
}

func main() {
	dflag := flag.Bool("d", false, "dump html")
	rflag := flag.Bool("r", false, "get rdf content")
	oflag := flag.String("o", ".", "set output directory")
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

