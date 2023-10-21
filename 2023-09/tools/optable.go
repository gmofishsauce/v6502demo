package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Names for all the node types that are supposed to be returned by the HTML parser
var typeNames = []string{"ErrorNode", "TextNode", "DocumentNode", "ElementNode", "CommentNode", "DoctypeNode"}

// An opFunc knows how to operate on one type of Node in an HTML doc.
// The Node may be e.g. "HTML comment" or "div tag".
type opFunc func(n *html.Node, cx *context) error

// We parse the document and make multiple passes over the resulting parse
// tree.  We have different operations (opFuncs) for the types in the tree
// per pass. The Golang HTML parser has Atoms to make lookups efficient,
// but it returns no atom for DOCTYPE nodes, comments, text nodes - only for
// HTML elements.  So an opTable needs to have operations for NodeTypes and
// also operations for specific tags. There is also a default action. When
// the action is nil, the processor will perform no operation.
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


