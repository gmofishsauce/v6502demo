package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type context struct {
	depth int
}

// An opFunc that returns an internal error
func internalError(n *html.Node, cx *context) error {
	return fmt.Errorf("internal error: no operation for node %v (context %v)", n, cx)
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

func getRDF(n *html.Node, cx *context) error {
	fmt.Printf("getRDF: attr=%v\n", n.Attr)
	return nil
}

var printPass = opTable{ printNode, [6]opFunc{}, nil, nil }

var rdfPass = opTable {
	defaultAction: nil,
	typeFuncs: [6]opFunc{},
	elementPreFuncs: map[atom.Atom]opFunc{atom.Link: getRDF},
	elementPostFuncs: nil,
}

func main() {
	in := os.Stdin
	name := "standard input"
	if len(os.Args) > 1 {
		name = os.Args[1]
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

	setOpTable(&rdfPass)
	if err = process(doc, &context{0}); err != nil {
		fmt.Fprintf(os.Stderr, "mkmd: process: %v\n", err)
		os.Exit(2)
	}

	fmt.Fprintf(os.Stderr, "mkmd: success\n")
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

// Markdown emitters
//
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
