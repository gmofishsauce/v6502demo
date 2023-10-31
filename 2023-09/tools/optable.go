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
