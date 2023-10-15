# Visual 6502 Capture and Save Project

## Overview

### Goal

Save as much of the original http://visual6502.org site as possible,
including the Wiki, the main site, and the images.

### What we have 

- The main site. On the main site, the Wiki is not accessible. Example: http://visual6502.org
- Wayback Machine backups of the Wiki. Example: https://web.archive.org/web/20210405071236/http://visual6502.org/wiki/index.php?title=Special%3AAllPages

## Steps

### Make a list of the main HTML pages

This was done by hand, by processing the index_php_title_Special_AllPages.html page
into a list of 75 HTML pages.

### Download the 75 pages

This was done with tools/getter.sh with the results stored in d_1_raw/

### Remove all the Wayback Machine crap

This was done by tools/fixer.sh which calls tools/fix_one.sh on each file.
This strips all the javascript and CSS links out of the header (in fact,
it strips out everything but the title tag) and strips the non-original
content from the body. Note: the original content is nicely delimited by
HTML comments in the Wayback Machine.

The results are stored in d_2_filtered/

### Get all the first-level images

This is done by tools/img_getter.sh which calls tools/get_img_for_file.sh
to get jpg, png, and gif extension files linked from the original 75 top
level HTML files. The results are placed in d_3_img/

Many of the results are HTML documents contained in files with one of the
three image file extensions. These files in turn contain more image links,
and they will need to be rewritten as markdown files for the eventual
md-based site. So these need to be "hoisted" - moved to be peers with the
original 75 files and any links to them rewritten to reference them.

Once they are hoisted, the images they link to can be fetched by rerunning
tools/img_getter.sh because it checks for the existence of a linked file
before hitting the Wayback Machine with wget. So it can be run as many
times as required after adding more HTML files.

### Hoist the images that are really HTML pages




