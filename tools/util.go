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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

    "golang.org/x/net/html"
)

// Context passed into the tree walker functions
type context struct {
	InputFilePath string
	OutputDirectory string
	NestingDepth int
	Markdown strings.Builder
	InFileToc bool
	InImgTag bool
	InJumpToNav bool
	InMagnify bool
	InFullImageLink bool
	InThumbCaption bool
	InATag bool
	ATagLeader string
	ATagContent string
	ATagTrailer string
	InMediaWikiFooter bool
	InWaybackMachineFooter bool
	InScript bool
	InTocNumber bool
	TableID string
	TableClass string

	listNestingDepth int
	tableColumnCount int
	tableNestingDepth int
	inTableHeader bool
}

func NewContext(outdir string, inpath string) *context {
	return &context{OutputDirectory: outdir, InputFilePath: inpath}
}

func (cx *context) EnterList() {
	cx.listNestingDepth++
}

func (cx *context) LeaveList() {
	cx.listNestingDepth--
}

func (cx *context) InList() bool {
	return cx.listNestingDepth > 0
}

// Enter a table. We suppress all nested tables in the wiki, so we
// we only to track one header state, that of the outermost table.
// We also suppress some outer tables.
func (cx *context) EnterTable() {
	cx.tableNestingDepth++
}

func (cx *context) LeaveTable() {
	cx.tableNestingDepth--
}

func (cx *context) EnterTableHeader() {
	cx.tableColumnCount = 0
	cx.inTableHeader = true
}

func (cx *context) AddTableColumn() {
	cx.tableColumnCount++
}

func (cx *context) GetTableColumns() int {
	return cx.tableColumnCount
}

func (cx *context) InTableHeader() bool {
	return cx.inTableHeader
}

func (cx *context) LeaveTableHeader() {
	cx.inTableHeader = false
}

func (cx *context) InTable() bool {
	return cx.tableNestingDepth > 0
}

// Nested tables are unconditionally ignored,
// meaning they generate no table markdown.
func (cx *context) InNestedTable() bool {
	return cx.tableNestingDepth > 1
}

// Suppressed tables generate no table markdown,
// but may have their content processed, and it
// may be processed specially (when compared to
// the same tags or content not in a suppressed
// table).
func (cx *context) InSuppressedTable() bool {
	if cx.TableClass == "toc" {
		return true
	}
	return false
}

// Emit a string to the standard output. The string should
// not contain any leading or trailing whitespace.
func (cx *context) emitString(content string) {
	if len(content) == 0 {
		return
	}
	if cx.InMediaWikiFooter {
		return
	}
	if cx.InJumpToNav || cx.InMagnify || cx.InThumbCaption {
		return
	}
	if cx.InNestedTable() {
		return
	}
	if cx.TableClass == "mw-allpages-table-form" {
		return
	}
	if cx.InTocNumber {
		return
	}
	if cx.InFileToc {
		return
	}

	// We always divert the entire contents of <A> tags EXCEPT for <IMG>
	// tags, which we want to emit directly. We emit the entire A tag
	// when we see the </A>. The effect is that when we have an IMG tag
	// wrapped in an A tag, we emit the entire IMG tag followed by the
	// entire A tag. Otherwise, we emit the A tag in-line. The caller,
	// of course, must clear InATag before trying to emit it, and must
	// set InImgTag before trying to emit that. I'm sure there's a better
	// way to deal with this.
	divert := cx.InATag && !cx.InImgTag 
	if divert {
		cx.ATagContent += content
	} else {
		cx.emitStringDirect(content)
	}
}

// This function writes to the markdown stream unconditionally. It can
// be used, with great care, to force  output when the surrounding tag
// is disabled.
func (cx *context) emitStringDirect(s string) {
	cx.Markdown.WriteString(s)
}

// ======= White space handling =======

// Newlines and spaces are absolutely critical in markdown; what's intuitive
// for human users is not obvious to a computer. We try to emit everything
// without any leading or trailing whitespace and instead insert these control
// characters which are ASCII DC1 through DC3 and so on if we need more.
const SingleSpace rune = 0x11 // ASCII DC1
const SingleNewline rune = 0x12
const DoubleNewline rune = 0x13 // ASCII DC3

func (cx *context) emitSingleSpaceNeeded() {
	cx.Markdown.WriteRune(SingleSpace)
}

func (cx *context) emitSingleNewlineNeeded() {
	cx.Markdown.WriteRune(SingleNewline)
}

func (cx *context) emitParagraphBreakNeeded() {
	cx.Markdown.WriteRune(DoubleNewline)
}

const nbsp = '\u00a0'

// If the argument string has leading spaces, tabs, or
// non-breaking space - but no newlines - return true.
func startsWithSpacesOnly(s string) bool {
	ok := false
	loop: // double break required in loop
	for _, c := range(s) {
		switch c {
		case ' ', '\t', nbsp:
			ok = true
		case '\n', '\r':
			ok = false
			break loop
		default:
			break loop
		}
	}
	return ok
}

func endsWithSpacesOnly(s string) bool {
	rs := []rune(s)
	ok := false
	loop: // double break required in loop
	for i := len(rs) - 1; i >= 0; i-- {
		switch rs[i] {
		case ' ', '\t', nbsp:
			ok = true
		case '\n', '\r':
			ok = false
			break loop
		default:
			break loop
		}
	}
	return ok
}

// Answer true if the rune is one of our "white markers"
// that represent white space.
func isWhiteMarker(r rune) bool {
	return r >= SingleSpace && r <= DoubleNewline
}

// Generate the final strings for a given white marker
// White markers must be combined before calling here;
// this function does not perform the combining.
func writeWhite(maxWhite rune, sb *strings.Builder) {
	switch maxWhite {
	case SingleSpace:
		sb.WriteRune(' ')
	case SingleNewline:
		sb.WriteRune('\n')
	case DoubleNewline:
		sb.WriteRune('\n')
		sb.WriteRune('\n')
	default:
		fatal("maxWhite")
	}
}

// Combine white markers in the argument string and
// return a new string with spaces and/or newlines.
func expandWhiteSpace(s string) string {
	var sb strings.Builder
	var inWhite bool
	var maxWhite rune

	for _, c := range s {
		if inWhite {
			// In a string of 1 or more whitemarker chars
			if isWhiteMarker(c) {
				// Keep the highest whitemarker in
				// the string of whitemarkers.
				if c > maxWhite {
					maxWhite = c
				} // else: skip lower or equal whitemarkers
			} else {
				// c is the first non-whitespace character
				// after substring of 1 or more whitemarkers
				writeWhite(maxWhite, &sb)
				sb.WriteRune(c)
				inWhite = false
			}
		} else {
			// Not in a string of whitemarkers
			if isWhiteMarker(c) {
				maxWhite = c
				inWhite = true
			} else {
				sb.WriteRune(c)
			}
		}
	}
	if inWhite {
		// the argument string s ended with whitemarkers
		writeWhite(maxWhite, &sb)
	}
	return sb.String()
}

func makeOutputDir(outputDirectory string) {
	err := os.MkdirAll(outputDirectory, 0700)	
	if err != nil && !errors.Is(err, fs.ErrExist) {
		fatal("unable to make output directory: %v\n", err)
	}
}

const waybackAPI = "https://archive.org/wayback/available?url="

func makeWaybackApiQuery(href string) string {
	return waybackAPI + makeAbsolute(href)
}

// Get the response body for the argument url. Return as a byte slice.
func getBody(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("getBody(): http.Get(%s): %v", url, err)
    }
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("getBody(): httpResponse.Read(%s): %v", url, err)
    }
	return b, nil
}

// Get the Wayback Machine URL of the latest snapshot corresponding to
// the href given by the link argument. Uses the WM's "available" API.
func getMostRecentUrl(url string) (string, error) {
	b, err := getBody(url)
	if err != nil {
		return "", err
	}

	var data map[string]any
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}

	archived_snapshots, ok := data["archived_snapshots"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("archived_snapshots not found in json response");
	}
	closest, ok := archived_snapshots["closest"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("closest not found in json response");
	}
	result, ok := closest["url"].(string)
	if !ok {
		return "", fmt.Errorf("url not found in json response");
	}
	return result, nil
}

// Return the title from the "title=" QS of the argument URL string.
// The result may be used as a filename, and we want to preserve the
// correspondence between filenames within the directory hierarchy.
// So we URL expand the string, converting e.g. %27 back to single
// quote, after which the caller can choose to run the URL-safe name
// corrector on the string.
func getTitle(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	parsedQuery, err := url.ParseQuery(parsedUrl.RawQuery)
	if err != nil {
		return "", err
	}
	titleString, ok := parsedQuery["title"]
	if !ok || len(titleString) != 1 {
		return "", fmt.Errorf("getTitle(%s): no title= query string", rawUrl)
	}
	result := titleString[0]
	return result, nil
}

// Return the value of the named attribute or ""
func getAttrVal(node *html.Node, name string) string {
	for _, v := range node.Attr {
		if v.Key == name {
			return v.Val
		}
	}
	return ""
}

// Return true if an attribute is present.
func hasAttr(node *html.Node, name string) bool {
	for _, v := range node.Attr {
		if v.Key == name {
			return true
		}
	}
	return false
}

// Make a safe URL from an URL that may contain illegal characters in
// the filename component only. Parse the URL, isolate the name from the
// path, unescape just the name, fix the name to contain no escapable
// characters, and put the URL back together - escaping it has been made
// unnecessary. This should only be applied to wiki URLs or else that
// last assumption may be incorrect.
func urlSafeUrl(origUrl string) string {
	// URLs external to the Wiki we don't want to mess with
	if strings.HasPrefix(origUrl, "http") {
		return origUrl
	}

	// Anchors we do a simple transformation on, which sometimes works.
	// Issue #011 points out that other times, it doesn't; needs work.
	if strings.HasPrefix(origUrl, "#") {
		// On Github anchors need to be #sym where sym is lower case
		// and separated by hyphens. In this wiki it seems like most
		// of them are mixed case, have no spaces and are separate by
		// underscores. This will handle them but not more.
		s := strings.ToLower(origUrl)
		s = strings.ReplaceAll(s, "_", "-")
		return s
	}

	u, err := url.Parse(origUrl)
	if err != nil {
		fatal("cannot parse URL %s: %v", origUrl, err)
	}

	// We need to convert the (last component of path name) +
	// (raw query string) to a combined filename with the invalid
	// URL character rule applied. There's no functionality on
	// the url package that does exactly what we need.
	s := u.Path
	m := u.Query()
	t, ok := m["title"]
	if ok {
		// Not worried about multiple "title=" query strings
		s += "?title=" + t[0] + ".md"
		if !strings.HasPrefix(s, "wiki") {
			s = "wiki/" + s
		}
	}
	dbg("u.Path=%s u.RawQuery=%s combined=%s", u.Path, u.RawQuery, s)
	s = makeUrlSafePath(expandEscapesInLine(s))

	// "If s doesn't start with prefix, s is returned unchanged."
	s = strings.TrimPrefix(s, "/wiki/")
	result := url.URL{
		Scheme: u.Scheme,
		Host: u.Host,
		Path: s,
	}
	dbg("urlSafeUrl %s => %s", origUrl, result.String())
	return result.String()
}

func isAsciiHexDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	if b >= 'A' && b <= 'F' {
		return true
	}
	// RFC recommends always using upper case
	// letters in %-encodings, but who knows.
	if b >= 'a' && b <= 'f' {
		return true
	}
	return false
}

func twoDigitHexAsciiToAsciiByte(ms byte, ls byte) byte {
	result := 16 * (ms - '0') // ms can't be > 7
	if ls <= '9' {
		result += ls - '0'
	} else if ls <= 'F' {
		result += ls - 'A'
	} else {
		result += ls - 'a'
	}
	return byte(result)
}

func expandEscapesInLine(s string) string {
	n := len(s)
	var sb strings.Builder
	skip := 0

	for i, c := range s {
		if skip > 0 {
			skip--
		} else if c == '%' && i+2 < n && isAsciiHexDigit(s[i+1]) && isAsciiHexDigit(s[i+2]) {
			sb.WriteByte(twoDigitHexAsciiToAsciiByte(s[i+1], s[i+2]))
			skip = 2
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func makeUrlSafePath(p string) string {
	dir := path.Dir(p)
	base := path.Base(p)
	result := path.Join(dir, urlSafeName(base))
	dbg("makeUrlSafePath %s => %s", p, result)
	return result
}

// This function is used to remap the name of a file to an URL-safe name.
// It must not be applied to a whole URL, because it will remap colons and
// slashes. The rule from RFC3986 is: ASCII letters and digits are legal,
// along with -_~. (hyphen, underscore, tilde, dot). Everything else is
// not legal in a name (the last component of a path). Many characters
// not legal in names are legal in query strings, i.e. encoding them is
// optional. The Wayback Machine Downloader makes query strings from the
// original MediaWiki into file names in the download, which is what causes
// the issue we're fixing.
func urlSafeName(origName string) string {
	var result strings.Builder
	for _, c := range origName {
		switch {
		case c >= 'a' && c <= 'z' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z':
			result.WriteRune(c)
		case c == '-' || c == '_' || c == '.' || c == '~':
			result.WriteRune(c)
		case c == '?' || c == '=' || c == ':':
			result.WriteByte('-')
		case c == '/':
			fatal("urlSafeName(): found '/' character: not a name: %s", origName)
		default:
			result.WriteByte('~')
		}
	}
	dbg("urlSafeName %s => %s", origName, result.String())
	return result.String()
}

func renameFileToUrlSafe(p string) error {
	safePath := makeUrlSafePath(p)
	dbg("Rename %s => %s", p, safePath)
	if err := os.Rename(p, safePath); err != nil {
		return err
	}
	return nil
}

func msg(format string, args... any) {
	var msg string
	if len(args) == 0 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, args)
	}
	fmt.Fprintf(os.Stderr, "mkmd: " + msg + "\n")
}

func warn(format string, args... any) {
	var msg string
	if len(args) == 0 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, args)
	}
	fmt.Fprintf(os.Stderr, "mkmd: warning: " + msg + "\n")
}

func fatal(format string, args... any) {
	var msg string
	if len(args) == 0 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, args)
	}
	fmt.Fprintf(os.Stderr, "mkmd: error: " + msg + "\nmkmd failed\n")
	os.Exit(1)
}

