# Log of actions and approaches

2023-09-06

Started. Wrote tools/getter.sh and over the next day, fixer.sh.
These are just quick hacks. I created a list of all the source
HTML pages ("the files") from the "all pages" link I got from
Ed:

https://web.archive.org/web/20210405071236/http://visual6502.org/wiki/index.php?title=Special%3AAllPages

Used to getter to download all the pages, but just the top level
HTML. The getter is not a crawler; it does not chase links.

2023-09-07

I chopped the Wayback Machine content out of the pages along with
all the style sheets and Javascript links. The files can be opened 
in a browser. The image links are all broken.

I wrote the basic framework of a recursive descent HTML parser for
translating the pages into markdown. It's a single file program in
tools/mkmd.go

2023-09-08

I figured out how thumbnails work and wrote some notes in image-notes.
I suspect that some of links were "Javascripted" and are now dead as
a result of stripping out the scripts. Since markdown processors generally
size images to fit, I'm thinking of generally writing tables to hold
the actual images (in the HTML, these held thumbnails) and not actually
capturing the thumbnails.  Whether I can do this depends on whether any
of the really huge images are in the wiki, or whether they are all on
the main site (the huge images are too big to fetch without an explicit
thumbnail click).

This needs more thought.

2023-09-09

The recursive descent parser (mkmd) should work on a page at a time and
should pull down the images, at least optionally, in addition to obviously
rewriting the links.

