package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Names for all the node types that are supposed to be returned by the HTML parser
var typeNames = []string{"ErrorNode", "TextNode", "DocumentNode", "ElementNode", "CommentNode", "DoctypeNode"}

// An opFunc knows how to operate on one type of Node in an HTML doc.
// The Node may be e.g. "HTML comment" or "div tag", i.e. it may be
// distinguished by NodeType or by both NodeType and Atoms (only when
// NodeType == ElementNode is there a nonzero Atom).
type opFunc func(n *html.Node, cx *context) error

// We parse the document and make multiple passes over the resulting parse
// tree.  We have different operations (opFuncs) for the types in the tree
// per pass. The Golang HTML parser has Atoms to make lookups efficient,
// but it returns no atom for DOCTYPE nodes, comments, text nodes - only for
// HTML elements.  So an opTable needs to have operations for NodeTypes and
// also operations for specific tags. There is also a default action. When
// any action or its entire containing table is nil, the processor for that
// pass will perform no operation.
type opTable struct {
	defaultAction opFunc
	typeFuncs [html.RawNode]opFunc
	elementPreFuncs map[atom.Atom]opFunc
	elementPostFuncs map[atom.Atom]opFunc
}

// The currentOpTable is set before each pass over the parse tree
var currentOpTable *opTable

// Set the opTable for the next pass
func setOpTable(c *opTable) {
	currentOpTable = c
}

func getOpTable() *opTable {
	return currentOpTable
}

// Return the operation.
func nodeToOp(n *html.Node, isPre bool) opFunc {
	var result opFunc
	if n.DataAtom != 0 {
		if isPre {
			result = currentOpTable.elementPreFuncs[n.DataAtom]
		} else {
			result = currentOpTable.elementPostFuncs[n.DataAtom]
		}
	} else {
		result = currentOpTable.typeFuncs[n.Type]
	}
	if result == nil {
		result = currentOpTable.defaultAction // may also be nil
	}
	return result
}

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
		atom.Atom]opFunc{atom.Link: rdfGetter},
	elementPostFuncs: nil,
}

