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

// Return the operation or nil if no operation should be performed on this
// Node at this time.
//
// There's a subtlety here: some Nodes can have a function attached to both
// the open tag and the closing tag. These are called the prefunc and the
// postfunc. But for nodes that don't have any kind of corresponding closing
// tag, like Doctype nodes and Text nodes, we only want to call one processing
// function on the Node (and we do it at prefunc time). At the same time, we
// want the default function, if any, to apply to any one of these calls. So
// the "if result == nil" clauses below cannot be factored out to the end of
// this function.
func nodeToOp(n *html.Node, isPre bool) opFunc {
	var result opFunc
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		if isPre {
			result = currentOpTable.elementPreFuncs[n.DataAtom]
		} else {
			result = currentOpTable.elementPostFuncs[n.DataAtom]
		}
		if result == nil {
			result = currentOpTable.defaultAction // may also be nil
		}
	} else { // DocType, TextNode, Comment (or Error)
		if isPre {
			result = currentOpTable.typeFuncs[n.Type]
			if result == nil {
				result = currentOpTable.defaultAction // may also be nil
			}
		}
	}
	return result
}

