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
	depth int
	outputDirectory string
	inputFilePath string
	md strings.Builder
	outputDisabledCounter int
}

func NewContext(outdir string, inpath string) *context {
	return &context{outputDirectory: outdir, inputFilePath: inpath}
}

// Emit a string to the standard output. For intended results (output)
// of this program. Avoid random fmt.PrintX to avoid random output.
func (cx *context) emitString(format string, args ...any) {
	if cx.outputDisabledCounter < 0 {
		panic("outputDisabledCounter negative")
	}
	if cx.outputDisabledCounter == 0 {
		cx.md.WriteString(fmt.Sprintf(format, args...))
	}
}

// Disable output generation, presumably until a matching end tag is found.
// This prevents e.g. emitting inline scripts which appear as text nodes.
func (cx *context) disableOutput() {
	cx.outputDisabledCounter++
}

func (cx *context) enableOutput() {
	cx.outputDisabledCounter--
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

// Implement -u. Fix names so they don't result in illegal URLs.

// This function is used to remap the name of a file to an URL-safe name.
// It must not be applied to a whole URL, because it will remap colons and
// slashes. The rule from RFC3986 is: ASCII letters and digits are legal,
// along with -_~. (dot). Everything else is not legal in a name (the last
// component of a path). Many characters not legal in names are legal in
// query strings, i.e. encoding them is optional. The Wayback Machine
// Downloader makes query strings from the original MediaWiki into file
// names in the download, which is what causes the issue we're fixing.
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

// Make a safe URL from an URL that may contain illegal characters in
// the filename component only. Parse the URL, isolate the name from the
// path, unescape just the name, fix the name to contain no escapable
// characters, and put the URL back together - escaping it has been made
// unnecessary. This should only be applied to wiki URLs or else that
// last assumption may be incorrect.
func urlSafeURL(origUrl string) string {
	return "FIXME"
}

func makeUrlSafeFileName(name string) error {
	dir := path.Dir(name)
	base := path.Base(name)
	safePath := path.Join(dir, urlSafeName(base))
	dbg("Rename %s => %s\n", name, safePath)
	if err := os.Rename(name, safePath); err != nil {
		return err
	}
	return nil
}

func fatal(format string, args... any) {
	var msg string
	if len(args) == 0 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, args)
	}
	fmt.Fprintf(os.Stderr, "mkmd: " + msg + "\n")
	os.Exit(1)
}

