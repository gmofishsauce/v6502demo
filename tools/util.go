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
	InHeader bool
	HeaderText string
	AnchorMap map[string]string
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
	return &context{OutputDirectory: outdir, InputFilePath: inpath, AnchorMap: map[string]string{}}
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
// characters which are ASCII DC1 through DC3 and so on.
//
// AnchorStart and AnchorEnd are used for a separate post-pass.
const SingleSpace rune = 0x11 // ASCII DC1
const SingleNewline rune = 0x12
const DoubleNewline rune = 0x13 // ASCII DC3
const AnchorStart rune = 0x1C // FS
const AnchorEnd rune = 0x1D // GS

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

// Combine white markers in the argument string and return
// a new string with spaces and/or newlines. This pass
// ignores anchors. See expandAnchors().
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

	// Anchors require special and elaborate handling.
	// See the comment above wikiMediaAnchor().
	if strings.HasPrefix(origUrl, "#") {
		return string(AnchorStart) + wikiMediaAnchor(origUrl[1:]) + string(AnchorEnd)
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
	}
	s = makeUrlSafePath(expandEscapesInLine(s, '%'))

	// "If s doesn't start with prefix, s is returned unchanged."
	s = strings.TrimPrefix(s, "/wiki/")
	result := url.URL{
		Scheme: u.Scheme,
		Host: u.Host,
		Path: s,
	}
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

func asciiRuneToTwoDigitHexAscii(c rune) string {
	const digs = "0123456789ABCDEF"
	return string(digs[(c >> 4)&0xF]) + string(digs[c&0xF])
}

// Expand '%27' sequences back to their non-URL equivalents.
// We also need to do this when the escape character is not
// the usual % sign - for explanation, see the long comment
// in urlSafeUrl().
func expandEscapesInLine(s string, escChar rune) string {
	n := len(s)
	var sb strings.Builder
	skip := 0

	for i, c := range s {
		if skip > 0 {
			skip--
		} else if c == escChar && i+2 < n && isAsciiHexDigit(s[i+1]) && isAsciiHexDigit(s[i+2]) {
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

// Anchor links. This is extremely complicated.
//
// Anchors are in-page links that target the "id" (or less commonly, "name")
// attributes of other tags. They are commonly used for table of contents
// links. The begin with a '#' character and do not contain a scheme or a
// path; they are just a string of legal URL characters (letters, digits,
// and - _ ~ . characters). This is all according to the specs.
//
// When the Jekyll Markdown processor reads the .md file and renders it as
// HTML, it automatically creates anchor targets ("id" attributes) for  <Hn>
// headings. It does this by applying some rules to the heading text. See:
//
// https://gist.github.com/asabaylus/3071099?permalink_comment_id=1593627#gistcomment-1593627
//
// 1. Jekyll downcases the heading (string).
// 2. It removes anything that is not a letter, number, space, hyphen, or
//		underscore (see the source via link above for how Unicode is handled)
// 3. It changes any space to a hyphen
// 4. If that is not unique, add "-1", "-2", "-3",... to make it unique
//
// Note that Github Markdown and Jekyll are not necessarily the same. The
// rules above originally referred to Github Markdown, and didn't cover the
// case of embedded _ (underscore) characters. I verified by eye that these
// rules seem for apply to Jekyll also, except that I added the case for
// underscores and didn't verify it for Github Markdown. Also, we don't bother
// even attempting to handle the last rule; hopefully this won't matter.
//
// But the WikiMedia software that generated the anchor links in the HTML we're
// reading also had a rule for processing headers into anchor links. Its rules
// were (apparently - I reverse engineered these) ...
//
// 1. Start with the header text (from the user's standpoint, the link target)
// 2. Remove all the spaces--leading, trailing, and embedded.
// 3. URL-escape the text, converting non-URL characters to %XX escapes.
// 4. Convert all non-URL characters, including % (percent) to dots.
//
// It doesn't _remove_ non-URL characters: it takes the escaped form of the URL
// with e.g. %27 and such, and simply replaces the % sign with a legal URL
// character, dot, to make e.g. .27 in the anchor URL.
//
// This transformation loses critical information - the word boundaries, which
// were preserved in Jekyll's transformation. So the WikiMedia anchor links are
// useless to us, although any anchor link that was just letters (e.g. "#title")
// makes it through both transformations unchanged and so happens to just work.
//
// So, what to do. There are several approaches that might work, but this one
// seems the most direct:
//
// 0. We assume that all anchor links target headers and don't handle others.
// 1. When we find a header (H1 through H6) we create an entry in global map.
//    The key is the header text with WikiMedia's algorithm applied. The value
//    is the header text with Jekyll's transformation applied.
// 2. When we find an anchor link, we write the link in the output, exactly as
//    we found it. But we surround it with ANCHOR_START and ANCHOR_END marker
//    characters (control characters).
// 3. When we expand whitespace at the end, we recognize the ANCHOR_START, get
//    the text to the ANCHOR_END, look the value up in the global map, and emit
//    value as the anchor link.
//
// Here's an example. The original page can be seen in the Wayback Machine here:
// https://web.archive.org/web/20210405071351/http://visual6502.org/wiki/index.php?title=RCA_1802E
//
// In the ToC, it says: 2.1 1. /ARO, /ALO
//
// This is made up a generated heading "2.1 " created by the processor for
// nested lists followed by a textual header "1. /ARO, /ALO". The "1. " in the
// textual header is probably a mistaken attempt by the original page author to
// give correct subheading numbers.
//
// The View Page source shows this for the link in the ToC:
// <li class="toclevel-2 tocsection-3">
//    <a href="#1._.2FARO.2C_.2FALO">
//        <span class="tocnumber">2.1</span>
//        <span class="toctext">1. /ARO, /ALO</span>
//    </a>
// </li>
//
// Below in the page, we find:
// <h4>
//     <span class="mw-headline" id="1._.2FARO.2C_.2FALO"> 1. /ARO, /ALO </span>
// </h4>
//
// When Jekyll processes the .md file we generate and renders it as HTML again,
// it generates the following HTML from the header tag:
//
// <h5 id="1-aro-alo">1. /ARO, /ALO</h5>
//
// This results from applying Jekyll's rules: (1) downcase the heading, giving
// "1. /aro, /alo"; (2) remove anything that is not a letter, number, space, hyphen,
// or underscore, giving "1 aro alo"; (3) replace spaces with hyphens, giving
// "1-aro-alo" as the id of the heading tag. This is what the anchor link must
// reference in the generated .md file.
//
// So when we encounter the ToC link, we write the href (which is the result of
// WikiMedia running its algorithm), surrounded by marker characters AnchorStart
// and AnchorEnd. (The emitted stream is buffered in memory for later postprocessing.)
//
// When we encounter the header, we run both algorithms and create a map from WikiMedia
// version of the anchor to the text of the Jekyll version. Finally, we post-process
// the emitted string looking for AnchorStart. When we find it, we replace the URL
// with the Jekyll version, which will work in the HTML rendered from the generated
// .md file. Hopefully.  ;-)
//

func isLetterOrDigit(c rune) bool {
	switch {
	case c >= 'a' && c <= 'z':
		return true
	case c >= 'A' && c <= 'Z':
		return true
	case c >= '0' && c <= '9':
		return true
	}
	return false
}

func isValidUrlCharNotTilde(c rune) bool {
	if isLetterOrDigit(c) {
		return true
	} else if c == '.' || c == '-' || c == '_' {
		return true
	}
	return false
}

// Convert an anchor link target like "1._/ARO,_/ALO" to an escaped
// fragment like "1._%2FARO%2C_%2FALO". The underscores in the argument
// string may come from the caller or from the original text being parsed.
func escapeFragment(f string) string {
	var sb strings.Builder
	for _, c := range f {
		if isValidUrlCharNotTilde(c) {
			sb.WriteRune(c)
		} else {
			sb.WriteString("%" + asciiRuneToTwoDigitHexAscii(c))
		}
	}
	return sb.String()
}

func wikiMediaAnchor(headerText string) string {
	// 1. Start with the header text (from the user's standpoint, the link target)
	// 2. Remove leading and trailing spaces; convert embedded spaces to underscores.
	// 3. URL-escape the text, converting non-URL characters to %XX escapes.
	// 4. Convert all non-URL characters, including % (percent) to dots.

	s := strings.ReplaceAll(strings.TrimSpace(headerText), " ", "_")
	f := escapeFragment(s)
	
	var sb strings.Builder
	for _, c := range f {
		if isValidUrlCharNotTilde(c) {
			sb.WriteRune(c)
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

func jekyllAnchor(headerText string) string {
	// 0. trim leading and trailing spaces - not documented but seems to be right
	// 1. Downcase the heading (string).
	// 2. Removes anything that is not a letter, number, space, hyphen, or	underscore
	// 3. Change any space to a hyphen

	var sb strings.Builder
	s := strings.TrimSpace(strings.ToLower(headerText))
	for _, c := range s {
		if c == ' ' {
			sb.WriteRune('-')
		} else if c == '_' {
			sb.WriteRune('_')
		} else if isLetterOrDigit(c) {
			sb.WriteRune(c)
		} // else nothing - drop c
	}
	return sb.String()
}

// Scan the entire text (generated markdown) for AnchorStart characters.
// When one is found, collect the text up AnchorEnd. Look the collected
// text up in the map and emit the value as the anchor.
func expandAnchors(s string, cx *context) string {
	var result strings.Builder
	var anchor strings.Builder
	inAnchor := false

	fmt.Fprintf(os.Stderr, "Anchor Map\n")
	for k, v := range cx.AnchorMap {
		fmt.Fprintf(os.Stderr, "Anchor Map \"%s\" => \"%s\"\n", k, v)
	}

	for _, c := range s {
		if !inAnchor {
			if c == AnchorStart {
				inAnchor = true
				anchor.Reset()
			} else {
				result.WriteRune(c)
			}
		} else { // in an anchor
			if c == AnchorEnd {
				jekyllAnchor, ok := cx.AnchorMap[anchor.String()]
				if !ok {
					jekyllAnchor = "link-could-not-be-patched"
				}
				result.WriteString("#" + jekyllAnchor)
				inAnchor = false
			} else {
				anchor.WriteRune(c)
			}
		}
	}
	return result.String()
}
