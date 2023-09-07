# Visual 6502 Capture and Save Project

## Overview

### Goal

Save as much of the original http://visual6502.org site as possible,
including the Wiki, the main site, and the images.

### What we have 

- The main site. On the main site, the Wiki is not accessible. Example: http://visual6502.org
- Wayback Machine backups of the Wiki. Example: https://web.archive.org/web/20210405071236/http://visual6502.org/wiki/index.php?title=Special%3AAllPages

## Strategy

### Download all the content

- Wiki: edit the All Pages page into script of wget commands
- Main site: maybe search for a crawler. Otherwise, do by hand.

### Clean up the HTML

All the HTML pages should be rewritten to contain bare HTML. This means nothing in the head section except the title tag. No headers or footers, particularly from the Wayback Machine.

This is straightforward for the Wiki pages. The other pages are TBD.

Need to think about links between pages. There are not a lot of these; the site structure is very simple.

Need to think about image links. TBD.

### Translate all the pages to markdown

There are a couple of tools that should be rechecked for usefulness after the HTML is cleaned: Cloudconvert and Codepen. Could also search for additional tools.

Otherwise, create a tool specific to the task.

### Create a static site that renders markdown

Needs a place to host the images. Mostly TBD.

### What's here

ALL_PAGES is a list of the 75 or so Wayback Machine URLs that make up the entire wiki.
The file was created by hand editing the file raw/index_php_title_Special_AllPages which
is a copy of the "All Pages" page from the Wiki on the Wayback Machine.

Directories:

#### raw

Download of the entire visual6502 wiki from the Wayback Machine

### content

Versions of every file in raw with the Wayback Machine headers and footers removed,
CSS removed, scripts removed, etc.

### log

Some notes about what was done at each step

### tools

Scripts and related hackery





