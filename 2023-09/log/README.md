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

2023-10-06

Today I cleaned up the image getter tools/img_getter.sh and successfully
downloaded all 35 images. The downoads in are rawimg/. The rawimg/ directory
is symlinked as img/ from the new md/ directory where the results will end up,
so images in md/ will only require "img/filename" as paths. The mkmd translator
doesn't know about this yet, so it's still putting broken web links in the
as-yet incomplete md-format output.

It looks like what's required now is just incremental improvement of mkmd and
then running it against all the files. I need to figure out how to bind Macdown
to the .md extension so I can just "open foo.md" and have Macdown open instead
of xcode. [Next day: done]

2023-10-07

Now that I understand these .jpg files that are actually HTML,the entire
processing pipeline is going to get more complicated. I should update the
Getter (tools/getter.sh) with the "&#160" fix in the Fixer (tools/fixer.sh)
so I can rerun the entire pipeline from scratch. (The very first thing in
the pipeline, the ALL_PAGES list, was made from the page
raw/index_php_title_Special_AllPages.html which I learned about from Ed.)

After the improved Getter runs and pulls all the pages in the ALL_PAGES
list into raw/, I need to run the img_getter to pull all the linked files
with .jpg and .png extensions into rawimg/.

Unfortunately, I now know that some of the ".jpg" files are actually HTML
documents with a bogus file extension. Handling for these is complicated.
I need to reprocess these with a tool like fixer that trims out the header
crap leaving just the title and the original body. Then rename them to
basename.html rather than .jpg and rewrite all the references to the file.

I probably need to leave them in the same directory because they contain
image references, although I can always rewrite anything to anything else
to fix that. Actually it might be easier to (1) move them to the content/
directory, (2) fix links to them to just be "./" file names, and (3) rewrite
their image links to "./img" (in the md/ directory). This makes the step of
fetching the images more consistent.

To do all this I should separate the loops over files from the processing
of the files. I need a script that loops building URLs from the ALL_PAGES
list and calls fix.sh on each one. The fixer 

So it's like this:

Pull all 75 pages named in ...Special_AllPages.html (Done)
Fix them up, like remove the XML-illegal entities (Done)
Fetch all the images referenced directly; some are really HTML (have tool)
These go in rawimg/ which is symlinked as md/raw (have tool)

For each .jpg that is really an HTML document:
  move the jpg to content/ and change its name to an HTML file
  Fix it up, strip out the header, fix entities, etc.
  Find all links to it (under its old name!) in the original 75 files
  And fix them

When done:
  Search all the N > 75 HTML files for any links to the original JPG names
  and fix them if there are any, etc.

