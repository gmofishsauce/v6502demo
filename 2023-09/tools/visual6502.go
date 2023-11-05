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
	"net/url"
	"regexp"
)

// Functions that are specific to the visual6502.org website

// This comment contains useful information, but no longer corresponds to the
// partitioning of the functionality in the program. The comment is about getting
// RDF files from the href in the `<link>` tag.
//
// The argument (href from the link tag) looks something like this:
// /wiki/index.php?title=6502DecimalMode&action=creativecommons
// This actually references a .rdf document with authorship information
//
// The Machine's API page for finding the most recent copy is:
// https://archive.org/wayback/available?url=
//
// An example of an URL that can be successfully looked up is:
// http://www.visual6502.org/wiki/index.php?title=6502DecimalMode&action=creativecommons
// ^^^^^^^^^^^^^^^^^^^^^^^^^ ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
// (this part is hardwired)  (url from href)
//
// ...although this worked, I think it should be URL encoded; also,
// the "action=creativecommons" is critical - that's what produces the RDF file.
//
// Sticking the two pieces together and visiting the API gives following result JSON:
//
//  {
//     "archived_snapshots" : {
//        "closest" : {
//           "available" : true,
//           "status" : "200",
//           "timestamp" : "20210405071434",
//           "url" : "http://web.archive.org/web/20210405071434/http://visual6502.org/wiki/index.php?title=6502DecimalMode"
//        }
//     },
//     "url" : "http://www.visual6502.org/wiki/index.php?title=6502DecimalMode"
//  }
//
// Note that the url returned in the result JSON **DOES NOT** produce the RDF page.
// It is necessary to manually edit the link two ways. First the `action=creativecommons`
// query string must be added, and then the magical `if_` must be added to the 14-character
// datetime substring. This will produce the RDF:
//
// wget -o wget.log -O rdf "http://web.archive.org/web/20210405071434if_/http://visual6502.org/wiki/index.php?title=6502DecimalMode&action=creativecommons"

/*
{
  "url": "http://visual6502.org/wiki/index.php?title=6502DecimalMode",
  "archived_snapshots":
  {
    "closest":
	{
	  "status": "200", "available": true,
	  "url": "http://web.archive.org/web/20210405071434/http://visual6502.org/wiki/index.php?title=6502DecimalMode",
	  "timestamp": "20210405071434"
	}
  }
}
*/
const wikiRoot =  "http://visual6502.org"

func makeAbsolute(wikiRef string) string {
	result, err := url.JoinPath(wikiRoot, wikiRef)
	if err != nil {
		fatal("could not create absolute path from site-relative href: %s %s\n", wikiRoot, wikiRef)
	}
	return result
}

const ccQS = "&action=creativecommons"
const magicIf = "if_"

// Fix up the URL of a top-level page so it references the page's licensing RDF
func fixupForRdf(url string) (string, error) {
	re := regexp.MustCompile(`.*[\d]{14}`)
	prefix := re.FindString(url)
	suffix := url[len(prefix):]
	if len(prefix) == 0 || len(suffix) == 0 {
		return "", fmt.Errorf("fixupForRdf(): failed to match URL")
	}
	return prefix+magicIf+suffix+ccQS, nil
}
