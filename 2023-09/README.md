# Visual 6502 Capture and Save Project

## Overview

### Goal

Save as much of the original http://visual6502.org site as possible,
including the Wiki, the main site, and the images.

### What we have 

- The main site. On the main site, the Wiki is not accessible. Example: http://visual6502.org
- Wayback Machine backups of the Wiki. Example: `https://web.archive.org/web/20210405071236/http://visual6502.org/wiki/index.php?title=Special%3AAllPages`

## Steps

### Make a list of the main HTML pages

This was done by hand, by processing the `index_php_title_Special_AllPages.html` page
into a list of 75 HTML pages.

### Download the 75 pages

This was done with `tools/getter.sh` with the results stored in `d_1_raw/`

### Remove all the Wayback Machine crap

This was done by `tools/fixer.sh` which calls `tools/fix_one.sh` on each file.
This strips all the javascript and CSS links out of the header (in fact,
it strips out everything but the title tag) and strips the non-original
content from the body. Note: the original content is nicely delimited by
HTML comments in the Wayback Machine.

The results are stored in `d_2_filtered/`

### Get all the first-level images

This is done by `tools/img_getter.sh` which calls tools/`get_img_for_file.sh`
to get jpg, png, and gif extension files linked from the original 75 top
level HTML files. The results are placed in `d_3_img/`

Many of the results are HTML documents contained in files with one of the
three image file extensions. These files in turn contain more image links,
and they will need to be rewritten as markdown files for the eventual
md-based site. So these need to be "hoisted" - moved to be peers with the
original 75 files and any links to them rewritten to reference them.

Once they are hoisted, the images they link to can be fetched by rerunning
`tools/img_getter.sh` because it checks for the existence of a linked file
before hitting the Wayback Machine with wget. So it can be run as many
times to download all the images as required after adding more HTML files.

### Hoist the images that are really HTML pages

Hoisting requires link rewriting which requires an understanding of links
in the wiki. I've identified the following types of links. This list isn't
known to be complete but the great majority of the links fit one of these.

1. Absolute links outside the wiki. In these links, the href starts with
`https://web.archive.org` and don't contain the string `visual6502.org/wiki`.
1. Absolute links within the wiki. These also start `https://web.archive.org`.
but do contain `visual6502.org/wiki`. I don't understand exactly why there
are such links because they somewhat defeat the purpose of having a Wiki.
1. Wiki links. These don't start with `https://web.archive.org` - they start
with the rest of link (`/web/...`) and contain `wiki/index.php?title=`
The target may an image, like `File:Atari_6507_7D.jpg` or it may be a page,
like `6502Observations`. I don't know if the title exactly corresponds to
the content of a title tag or is processed (e.g. spaces removed).
1. Anchor links. Here the href is e.g. `href="#Tests_for_ADC"`. These are
of course in-page links.
1. Certain `<link>` tags in the header. See below.

I think the external absolute links should be left as is, pointing into
the Wayback Machine for now, at last until we host the new Wiki somewhere
else than Github. This avoids licensing questions about the external
images.

Wiki-internal absolute links will be rewritten to point to their corresponding `.md` files in the new Wiki. This includes image links within the Wiki. Anchors will be handled as best they can be in Markdown. Initial reading is encouraging.

2023-10-16

There's another kind of link: <link> tags in the header of the basic 75
files point to .rdf files that contain authorship information. I wasn't
aware because `fixer.sh` strips everything from the header except the
`<title>` tag. The idea behind this was to get rid of the Javascript and
CSS links. Correcting this oversight requires modifying the Fixer and
regenerating all the files in d_2_filtered.

Other interesting behaviors may be visible as a result of perusing the
list of all the files captured by the Wayback Machine. The Wayback Machine
considers every combination of query string arguments to be separate links.
Unfortunately this means is captures thouands and thousands of copies of
the "most recent updates" page for various numbers of days. But more
generally it's useful in figuring out the behavior of the Wiki's Javascript.

I'm considering giving up on most of the downloader scripts and writing a
purpose-built crawler in Golang. More generally I need to design the layout
of the eventual markdown site in order to figure out how to rewrite the
links in the downloaded documents.



