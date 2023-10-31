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
	"os"
	"regexp"
)

func makeOutputDir(outputDirectory string) {
	err := os.MkdirAll(outputDirectory, 0700)	
	if err != nil && !errors.Is(err, fs.ErrExist) {
		fatal("unable to make output directory: %v\n", err)
	}
}

// XXX this url has the query string on the end. See TODO below.
const waybackAPI = "https://archive.org/wayback/available?url="

func makeWaybackApiQuery(href string) string {
	// TODO figure out how to use the url package to add a query string.
	return waybackAPI + makeAbsolute(href)
}

// Get the response body for the argument url. Return as a byte slice.
func getBody(url string) ([]byte, error) {
	dbg("getBody(%s)\n", url)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("getBody(): http.Get(%s): %v", url, err)
    }
	defer resp.Body.Close()
    dbg("getBody() resp: %v\n", resp)

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
	dbg("unmarshaled response: %v\n", data)

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

func getTitle(url string) (string, error) {
	prefix := regexp.MustCompile(`.*title=`).FindString(url)
	suffix := regexp.MustCompile(`&.*$`).FindString(url)
	if len(prefix) == 0 {
		return "", fmt.Errorf("getTitle(): failed to match URL")
	}
	return url[len(prefix):len(url)-len(suffix)], nil
}

func fatal(format string, args... any) {
	msg := fmt.Sprintf(format, args)
	fmt.Fprintf(os.Stderr, "mkmd: " + msg)
	os.Exit(1)
}


