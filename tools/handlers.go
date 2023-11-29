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
	typeFuncs: [6]opFunc{nil, doText, nil, nil, doComment, doDocType},
	elementPreFuncs: map[atom.Atom]opFunc{
		atom.A:  doAtagOpen,
		atom.Div: doDivOpen,
		atom.H1: doHeaderOpen,
		atom.H2: doHeaderOpen,
		atom.H3: doHeaderOpen,
		atom.H4: doHeaderOpen,
		atom.H5: doHeaderOpen,
		atom.H6: doHeaderOpen,
		atom.Img: doImgOpen,
		atom.Li: doLiOpen,
		atom.Ol: doOlOpen,
		atom.P:  doPtagOpen,
		atom.Pre: doPreOpen,
		atom.Script: doScriptOpen,
		atom.Span: doSpanOpen,
		atom.Table: doTableOpen,
		atom.Title: doTitleOpen,
		atom.Td: doTdOpen,
		atom.Th: doThOpen,
		atom.Tr: doTrOpen,
		atom.Ul: doUlOpen,
	},
	elementPostFuncs: map[atom.Atom] opFunc{
		atom.A:  doAtagClose,
		atom.Body: doHtmlClose,
		atom.Div: doDivClose,
		atom.H1: doHeaderClose,
		atom.H2: doHeaderClose,
		atom.H3: doHeaderClose,
		atom.H4: doHeaderClose,
		atom.H5: doHeaderClose,
		atom.H6: doHeaderClose,
		atom.Html: doHtmlClose,
		atom.Li: doLiClose,
		atom.Ol: doOlClose,
		atom.P:  doPtagClose,
		atom.Pre: doPreClose,
		atom.Script: doScriptClose,
		atom.Span: doSpanClose,
		atom.Table: doTableClose,
		atom.Title: doTitleClose,
		atom.Td: doTdClose,
		atom.Th: doThClose,
		atom.Tr: doTrClose,
		atom.Ul: doUlClose,
	},
}

// An optable for dumping the entire document for debugging purposes
var xPass = opTable{
	defaultAction: nil,
	typeFuncs: [6]opFunc{},
	elementPreFuncs: map[atom.Atom] opFunc{
		atom.Table: xFuncOpen,
	},
	elementPostFuncs: map[atom.Atom] opFunc{
		atom.Table: xFuncClose,
	},
}

// =============================================================
// Operation functions
// =============================================================

func prototype(n *html.Node, cx *context) error {
	return nil // copy this to make new handlers
}

/*

This is an example of a reoccurring block that I need to learn how to treat.
It's got three divs with recognizable classes containing an A tag that links
to a description page (an HTML page despite the name that ends with .png).

Then there's an img tag in the A tag with a link to a legitimate image.

<div class="center">
    <div class="thumb tnone">
        <div class="thumbinner" style="width:280px;">
            <a href="/wiki/index.php?title=File:NES-2A03-decimal-DAA-removed.png" class="image">
                <img alt="" src="/wiki/images/8/89/NES-2A03-decimal-DAA-removed.png" width="278" height="200" class="thumbimage" />
            </a>
            <div class="thumbcaption">
                Transistor t2556 in NES 2A03
            </div>
        </div>
    </div>
</div>

There's a similar bunch of code distinguished by a link to the currently
non-existent (in the wiki target directory) url skins/common/images/magnify-clip.png
that also needs to be addressed.

*/
// According to the standard, A-tags don't nest
func doAtagOpen(n *html.Node, cx *context) error {
	if cx.InATag {
		fatal("nested A-tag")
	}

	// MediaWikis have a bunch of HTML files with names like foo.png
	// and bar.jpg that are actually HTML files which wrap a thumb
	// and have a link to a larger version of the image. We don't
	// output these.
	cl := getAttrVal(n, "class")
	if cl == "image" {
		return nil
	}
	href := getAttrVal(n, "href")
	if len(href) == 0 {
		warn("A-tag with no href")
		return nil
	}
	if strings.HasSuffix(href, ".gif") || strings.HasSuffix(href, ".png") || strings.HasSuffix(href, ".jpg") {
		// This is a link surrounding a thumbnail. We need to find
		// a solution for the (few) cases where "larger image" is
		// actually larger than the thumbnail.
		return nil
	}

	cx.InATag = true
	cx.emitString("[")
	cx.ATagRemainder = "](" + urlSafeUrl(href) + ")"
	return nil
}

func doAtagClose(n *html.Node, cx *context) error {
	// See comment above
	cl := getAttrVal(n, "class")
	if cl == "image" {
		return nil
	}
	href := getAttrVal(n, "href")
	if strings.HasSuffix(href, ".gif") {
		return nil
	}

	if len(cx.ATagRemainder) == 0 {
		warn("/A-tag with no remainder")
		return nil
	}
	cx.emitString(cx.ATagRemainder)
	cx.InATag = false
	cx.ATagRemainder = ""
	return nil
}

/*

Chop out some divs. Example: MediaWiki navigation links.

Need to deal with this: we don't want the magnify-clip.png,
but do want the caption, "Transistor t2556 in visual6502" in this case.
I found 38 occurrences of this pattern in the wiki.
We suppress class="thumbcaption" and instead emit a newline,
which causes the caption to be left-justified.

<div class="center">
    <div class="thumb tnone">
        <div class="thumbinner" style="width:280px;">
            <a href="/wiki/index.php?title=File:6502-decimal-DAA-removed-visual6502.png" class="image">
            <img alt="" src="/wiki/images/thumb/7/78/6502-decimal-DAA-removed-visual6502.png/278px-6502-decimal-DAA-removed-visual6502.png" width="278" height="194" class="thumbimage" />
            </a>
            <div class="thumbcaption">
                <div class="magnify">
                    <a href="/wiki/index.php?title=File:6502-decimal-DAA-removed-visual6502.png" class="internal" title="Enlarge">
                        <img src="/wiki/skins/common/images/magnify-clip.png" width="15" height="11" alt="" />
                    </a>
                </div>
                Transistor t2556 in visual6502
            </div>
        </div>
    </div>
</div>

*/

func doDivOpen(n *html.Node, cx *context) error {
	id := getAttrVal(n, "id")
	if id == "jump-to-nav" {
		cx.InJumpToNav = true
		return nil
	}
	cl := getAttrVal(n, "class")
	if cl == "magnify" {
		cx.InMagnify = true
		return nil
	}
	if cl == "fullImageLink" {
		cx.InFullImageLink = true
		return nil
	}
	if cl == "thumbcaption" {
		cx.InThumbCaption = true
		return nil
	}

	return nil
}

func doDivClose(n *html.Node, cx *context) error {
	id := getAttrVal(n, "id")
	if id == "jump-to-nav" {
		cx.InJumpToNav = false
		return nil
	}
	cl := getAttrVal(n, "class")
	if cl == "magnify" {
		cx.InMagnify = false
		return nil
	}
	if cl == "fullImageLink" {
		cx.InFullImageLink = true
		return nil
	}
	if cl == "thumbcaption" {
		cx.InThumbCaption = false
		return nil
	}

	return nil
}

// </html> endtag. Write the output file.
func doHtmlClose(n *html.Node, cx *context) error {
	outDir := cx.OutputDirectory
	if len(outDir) == 0 {
		outDir = "."
	}
	inFileName := path.Base(cx.InputFilePath)
	outPath := path.Join(outDir, inFileName) + ".md"
	s := expandWhiteSpace(cx.Markdown.String())
	err := os.WriteFile(outPath, []byte(s), 0644)
	if err != nil {
		fatal("write result to \"%s\" failed: %v", outPath, err)
	}
	return nil
}

func doImgOpen(n *html.Node, cx *context) error {
	imgText := getAttrVal(n, "alt")
	if len(imgText) == 0 {
		imgText = "Image (no description given)"
	}

	imgLink := getAttrVal(n, "src")
	if len(imgLink) == 0 {
		warn("img tag with no src")
		return nil
	}
	if strings.Contains(imgLink, "skins/common") {
		// "Powered by MediaWiki" image
		return nil
	}

	// Make the image addressable. Note: comment from docs:
	// "If s doesn't start with prefix, s is returned unchanged."
	imgLink = strings.TrimPrefix(imgLink, "/wiki/")
	cx.emitString("\n\n![%s](%s)\n\n", imgText, imgLink)
	return nil
}

func doHeaderOpen(n *html.Node, cx *context) error {
	if cx.InTable() {
		// Markdown can't put headers in table cells,
		// we just turn the header into boldface text.
		cx.emitSingleSpaceNeeded()
		cx.emitString("**")
		return nil
	}
	const hashes = "#######"
	cx.emitParagraphBreakNeeded()
	cx.emitString(hashes[0:1 + int(n.Data[1] - '0')])
	cx.emitSingleSpaceNeeded()
	return nil
}

func doHeaderClose(n *html.Node, cx *context) error {
	if cx.InTable() {
		cx.emitString("**")
		cx.emitSingleSpaceNeeded()
		return nil
	}
	cx.emitSingleNewlineNeeded()
	return nil
}

// See issue #006 and #007.
func doLiOpen(n *html.Node, cx *context) error {
	cx.emitSingleNewlineNeeded()
	cx.emitString("-")
	cx.emitSingleSpaceNeeded()
	return nil
}

func doLiClose(n *html.Node, cx *context) error {
	return nil
}

func doOlOpen(n *html.Node, cx *context) error {
	if cx.InList() {
		// entering a nested list
		cx.emitSingleNewlineNeeded()
	} else {
		cx.emitParagraphBreakNeeded()
	}
	cx.EnterList()
	return nil
}

func doOlClose(n *html.Node, cx *context) error {
	cx.LeaveList()
	if cx.InList() {
		cx.emitSingleNewlineNeeded()
	} else {
		cx.emitParagraphBreakNeeded()
	}
	return nil
}

func doPtagOpen(n *html.Node, cx *context) error {
	cx.emitParagraphBreakNeeded()
	return nil
}

func doPtagClose(n *html.Node, cx *context) error {
	cx.emitParagraphBreakNeeded()
	return nil
}

func doPreOpen(n *html.Node, cx *context) error {
	cx.emitSingleNewlineNeeded()
	cx.emitString("```")
	cx.emitSingleNewlineNeeded()
	return nil
}

func doPreClose(n *html.Node, cx *context) error {
	cx.emitSingleNewlineNeeded()
	cx.emitString("```")
	cx.emitSingleNewlineNeeded()
	return nil
}

func doSpanOpen(n *html.Node, cx *context) error {
	cl := getAttrVal(n, "class")
	if cl == "tocnumber" {
		cx.InTocNumber = true
	}
	return nil
}

func doSpanClose(n *html.Node, cx *context) error {
	cl := getAttrVal(n, "class")
	if cl == "tocnumber" {
		cx.InTocNumber = false
	}
	return nil
}

/*

This was seen in work/wiki/index.php-title-6502_Opcode_8B_~XAA~_ANE~
It's possible we should just suppress this table.
I think it exists to make the whole TOC accessible
to a piece of Javascript that can disable it.

<table id="toc" class="toc">
    <tr>
        <td>
            <div id="toctitle">
                <h2>Contents</h2>
            </div>
            <ul>
                <li class="toclevel-1 tocsection-1"><a href="#Explanation">
                    <span class="tocnumber">1</span> <span class="toctext">Explanation</span></a></li> <li class="toclevel-1 tocsection-2"><a href="#Circuit_Diagram">
                    <span class="tocnumber">2</span> <span class="toctext">Circuit Diagram</span></a></li>
                <li class="toclevel-1 tocsection-3"><a href="#Testing_this_opcode">
                    <span class="tocnumber">3</span> <span class="toctext">Testing this opcode</span></a></li>
                <li class="toclevel-1 tocsection-4"><a href="#Modelling_this_opcode">
                    <span class="tocnumber">4</span> <span class="toctext">Modelling this opcode</span></a></li>
                <li class="toclevel-1 tocsection-5"><a href="#Tested_CPUs">
                    <span class="tocnumber">5</span> <span class="toctext">Tested CPUs</span></a></li>
                <li class="toclevel-1 tocsection-6"><a href="#Resources">
                    <span class="tocnumber">6</span> <span class="toctext">Resources</span></a></li>
            </ul>
        </td>
    </tr>
</table>
*/

// Table processing - the trickiest part of the whole business.
// Immediately enter the table, because we need to track nesting.
// But if entering has put us in a nested table, we're done - we
// completely suppress table output from all nested tables (this
// turns out not to lose any content in this wiki, just controls
// that wouldn't work anyway). Next grab the id and class from
// this top-level table into the context. Now it's possible to
// check if table-tag processing for this table type is suppressed.
// This only means we don't generate table markdown: we may still
// process and generate output for content in specific types of
// "suppressed" tables on a case by case basis. Finally, if this
// table isn't nested and isn't suppressed, begin generating the
// markdown for a table, starting with a paragraph break. This
// processing outline necessarily applies to all the table tags.
func doTableOpen(n *html.Node, cx *context) error {
	cx.EnterTable()
	if cx.InNestedTable() {
		return nil
	}

	cx.TableID = getAttrVal(n, "id")
	cx.TableClass = getAttrVal(n, "class")
	if cx.InSuppressedTable() {
		return nil
	}

	cx.EnterTableHeader()
	cx.emitParagraphBreakNeeded()
	return nil
}

func doTableClose(n *html.Node, cx *context) error {
	cx.LeaveTable()
	if cx.InNestedTable() {
		return nil
	}

	if !cx.InSuppressedTable() { // ???
		cx.emitParagraphBreakNeeded()
	}
	cx.TableID = ""
	cx.TableClass = ""
	return nil
}

func doTdOpen(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	if cx.InTableHeader() {
		cx.AddTableColumn()
	}
	cx.emitString("|")
	cx.emitSingleSpaceNeeded()
	return nil
}

func doTdClose(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	cx.emitSingleSpaceNeeded()
	return nil
}

func doThOpen(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	if cx.InTableHeader() {
		cx.AddTableColumn()
		cx.emitString("|")
		cx.emitSingleSpaceNeeded()
	}
	return nil
}

func doThClose(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	cx.emitSingleSpaceNeeded()
	return nil
}

func doTrOpen(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	cx.emitSingleNewlineNeeded()
	return nil
}

func doTrClose(n *html.Node, cx *context) error {
	if cx.InNestedTable() || cx.InSuppressedTable() {
		return nil
	}
	cx.emitString("|")
	cx.emitSingleNewlineNeeded()

	if cx.InTableHeader() {
		for i := 0; i < cx.GetTableColumns(); i++ {
			cx.emitString("|:---:")
		}
		cx.emitString("|")
		cx.emitSingleNewlineNeeded()
		cx.LeaveTableHeader()
	}
	return nil
}

func doTitleOpen(n *html.Node, cx *context) error {
	cx.emitSingleNewlineNeeded()
	cx.emitString("#")
	cx.emitSingleSpaceNeeded()
	return nil
}

func doTitleClose(n *html.Node, cx *context) error {
	cx.emitSingleNewlineNeeded()
	return nil
}

func doUlOpen(n *html.Node, cx *context) error {
	if cx.InList() {
		// entering a nested list
		cx.emitSingleNewlineNeeded()
	} else {
		// entering a first level list
		cx.emitParagraphBreakNeeded()
	}
	cx.EnterList()
	return nil
}

func doUlClose(n *html.Node, cx *context) error {
	cx.LeaveList()
	if cx.InList() {
		// still in a first (n-1st) level list
		cx.emitSingleNewlineNeeded()
	} else {
		// exited a first level list
		cx.emitParagraphBreakNeeded()
	}
	return nil
}

func doDocType(n *html.Node, cx *context) error {
	cx.emitString("**INCOMPLETE DRAFT OF RECOVERED WIKI PAGE**\n")
	return nil
}

func doText(n *html.Node, cx *context) error {
	if cx.InScript {
		return nil
	}
	s := strings.TrimSpace(n.Data)
	if len(s) != 0 {
		if startsWithSpacesOnly(n.Data) {
			cx.emitSingleSpaceNeeded()
		}
		cx.emitString(strings.ReplaceAll(s, "_", "\\_"))
		if endsWithSpacesOnly(n.Data) {
			cx.emitSingleSpaceNeeded()
		}
	}
	return nil
}

func doComment(n *html.Node, cx *context) error {
	if strings.Contains(n.Data, "end content") {
		cx.InMediaWikiFooter = true // continues to end
	}
	if strings.Contains(n.Data, "Saved in parser cache") {
		cx.InWaybackMachineFooter = true // continues to end
	}
	return nil
}

func doScriptOpen(n *html.Node, cx *context) error {
	cx.InScript = true
	return nil
}

func doScriptClose(n *html.Node, cx *context) error {
	cx.InScript = false
	return nil
}

// End of support for -m. Begin support for -r.

// rdfHandler is invoked for <link> tags only. It identifies tags
// that link .rdf files containing authorship and license data and
// downloads their latest version.
func rdfHandler(n *html.Node, cx *context) error {
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

	// If we have the correct kind of link, fetch it. Note: We want to
	// maintain certain correspondences between file names. But the Jekyll
	// processor (or some part of the Gitlab Pages pipeline) is not happy
	// with filenames having illegal URL characters like single quotes or
	// parens, and three of the top level files (and their corresponding
	// .rdf files) contain single quotes. Yet all we have here is the href
	// to the file, which does not contain a single quote - it contains an
	// URL-encoded single quote, %27. Eventually I made the decision to
	// create a substitution rule and rename all the files. This means I
	// need to make the same substitutions in the URLs that refer to them.
	// So I need to get the filename component of the URL with the %27
	// expanded to a single quote, then run the substitution rule on the
	// the resulting string, then write the file under that name, and then
	// finally (later, not here) fix the href to match. The substitution
	// rule for characters produces only URL-safe characters, so there is
	// never any need to URL encode after running the rule.
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
		rdfName = urlSafeName(rdfName)
		makeOutputDir(cx.OutputDirectory)
		if err = os.WriteFile(path.Join(cx.OutputDirectory, rdfName + ".rdf"), page, 0600); err != nil {
			return err
		}
	}

	return nil
}

// End of support for -r. Begin general handlers.

// An opFunc that returns an internal error
func internalError(n *html.Node, cx *context) error {
	return fmt.Errorf("internal error: type %v not expected (context %v)", n, cx)
}

// An opFunc that prints "not handled: thing" for use as a default
func notHandled(n *html.Node, cx *context) error {
	dbg("not handled: node %v (context %v)", n, cx)
	return nil
}

// A debugging opFunc that just prints the node with indent
func printNode(n *html.Node, cx *context) error {
	fmt.Printf("%*sType=%s DataAtom=%v Data=%v Attr=%v\n", cx.NestingDepth*2, "",
		typeNames[n.Type], n.DataAtom, strings.TrimSpace(n.Data), n.Attr)
	return nil
}

var tableDepth int = 0 // special purpose hack, don't put in context

func xFuncOpen(n *html.Node, cx *context) error {
	if n.DataAtom == atom.Table {
		tableDepth++
		if tableDepth > 1 {
			fmt.Printf("%s: found table depth > 1\n", cx.InputFilePath)
		}
	}
    return nil
}

func xFuncClose(n *html.Node, cx *context) error {
	if n.DataAtom == atom.Table {
		tableDepth--
	}
    return nil
}
