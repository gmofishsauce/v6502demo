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

This supposedly works; from
https://stackoverflow.com/questions/56883209/thumbnails-of-images-in-github-markdown

You can add a smaller version of the image with a link to the full resolution one.
[![sample screenshot](https://i.imgur.com/Tkks00R.png)](https://i.imgur.com/Ob4qAwu.png)

2023-09-10

I figured out that you need to read at least an entire row of an HTML table before
you know how many columns there are. Only then can you start emitting the markdown
header for the table. This requires "infinite" lookahead, aka "two passes". Since
I want this tool to work in a pipeline and I can't actually reread the standard
input, I'll have to create a list of all the tokens that I can reread and hide
the list behind a getToken() kind of call. This front end reader can make copies
of every token so the rest of the tool can do whatever it wants with them.

2023-09-11

After significant revisions last night, I'm closer to a pure recurive descent
parser that handles alternatives. Among other changes, I implemented one token
of pushback. This allows each parsing function in the recursive descent to handle
both its own end tag and its own start tag, because the parent pushes back the
start tag before invoking the function for it. This is significant because the
start tag may have attributes that are important to generated output.

The next step is to generalize the notion of an expected set of tags. At the
outer level of the body, the expected set would be large. Inside a table row,
the expected set might be just "td". 

2023-09-15

I've been intermittently working to improve the HTML parser in mkmd, not yet
attempting to either build a tree or generate any output except a few samples.

Today I discovered that HTML entities like &nbsp; are not legal XML, which puts
the lie to the concept of "XHTML" being an XML language. I modified the "Fixer"
to translate &nbsp; to &#160; which is legal XML.


