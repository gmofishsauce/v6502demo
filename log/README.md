# Introduction

I ended up with three README files that each contained bits and
pieces of the history of the work here. On 2023/10/27, I combined
the README from ".." into this file. The end of the incorporated
README is marked.


## Visual 6502 Capture and Save Project

### Overview

#### Goal

Save as much of the original http://visual6502.org site as possible,
including the Wiki, the main site, and the images.

#### What we have 

- The main site. On the main site, the Wiki is not accessible. Example: http://visual6502.org
- Wayback Machine backups of the Wiki. Example: `https://web.archive.org/web/20210405071236/http://visual6502.org/wiki/index.php?title=Special%3AAllPages`

### Steps

#### Make a list of the main HTML pages

This was done by hand, by processing the `index_php_title_Special_AllPages.html` page
into a list of 75 HTML pages.

#### Download the 75 pages

This was done with `tools/getter.sh` with the results stored in `d_1_raw/`

#### Remove all the Wayback Machine crap

This was done by `tools/fixer.sh` which calls `tools/fix_one.sh` on each file.
This strips all the javascript and CSS links out of the header (in fact,
it strips out everything but the title tag) and strips the non-original
content from the body. Note: the original content is nicely delimited by
HTML comments in the Wayback Machine.

The results are stored in `d_2_filtered/`

#### Get all the first-level images

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

#### Hoist the images that are really HTML pages

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
regenerating all the files in `d_2_filtered`.

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

## END

That's the end of the old "../README.md" file, added here on 10/27.

The rest of this document is more or less a log.

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

Today I cleaned up the image getter `tools/img_getter.sh` and successfully
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
the pipeline, the `ALL_PAGES` list, was made from the page
`raw/index_php_title_Special_AllPages.html` which I learned about from Ed.)

After the improved Getter runs and pulls all the pages in the `ALL_PAGES`
list into raw/, I need to run the `img_getter` to pull all the linked files with .jpg and .png extensions into rawimg/.

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
of the files. I need a script that loops building URLs from the `ALL_PAGES`
list and calls fix.sh on each one. The fixer 

So it's like this:

Pull all 75 pages named in `...Special_AllPages.html` (Done)
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

### 2023/10/27

Todat I merged ../README.md into this file. I also located all the symbols with underscores and wrapped them in single back quotes so they'll render correctly.

During the previous two weeks, I backed off almost completely from using scripts and decided to write a downloader and reformatter in Golang. The only thing I kept from the scripting days was `d_1_raw/`, the 75 or so top-level files that are helpful for debugging.

Then I began working to download all the RDF files with the authorship information. Eventually I got stuck (on an issue that turned out to be simple and stupid; but in such an unfamiliar world technically, given my lack of web and Javascript experience, it was hard to troubleshoot.)

I posed a question on the ex-New Relic Slack and got some interesting suggestions.
One was a tool written in Ruby, a gem that can execute from the command line
as [`wayback_machine_downloader`](https://github.com/hartator/wayback-machine-downloader).

I did a bunch of experiments with it. Of course initially it wanted to download 28,000 files for the Wiki alone, because the WM has saved that many (one for every value of "most recent=N days" that appears anywhere in a link, etc.) I eventually ended up with

```
wayback_machine_downloader http://visual6502.org --only wiki \
  --exclude "/\&[A-Za-z]+|Special:/" > all_files.wmd 2>stderr.wmd
```

Yes, the log files are WMDs.  ;-)

You can apparently only have one `--exclude` option; they do not seem to accumulate. So the regex has get more and more complex instead, using pipe characters to separate alternatives.

This seems to have downloaded maybe 750 or so files, which is possible; it's about 10 images and secondary files for each top-level page. I'm currently investigating what I've got and in what format.

I moved the downloaded `websites/` directory to 2023-09, got rid of the `d_N...` directories (actually saved in `~/Fuss/V6502WikiOldDownloads`) and dumped the log files, which were mostly lists of the files that did not match the patterns.

I found that some of the downloaded files were gzipped and then I found that gzip requires the .gz file extension, which wasn't present. I renamed the top level gzipped files:

```
file * | grep 'gzip compressed data, from Unix' | while read i ; do name=$(echo "$i" | sed 's/ .*$//' | sed 's/:$//') ; mv "$name" "${name}.gz" ; done
```

2023/10/28

I committed the changes above. Now it turns out that the WMD doesn't seem to know anything about the .rdf documents which contain the authorship information.
These are referenced only from `<link>` tags so it's not surprising. But it means my Golang downloader code was **not** a waste of time.
Going back to work on that.

2023/10/29

I got the RDF downloader working! But I need to decide where to put the RDF files,
which is similar to thinking through the layout of the entire eventual markdown-based wiki.

I found four more compressed files, using a complete search that handles the one filename
that has a single quote in it:

```
Macbook-Pro-2019:2023-09 jeff$ find . -type f -print0 | xargs -0 file | grep gzip

./websites/visual6502.org/wiki/skins/common/diff.css?270: gzip compressed data...
./websites/visual6502.org/wiki/skins/common/diff.js?270:  gzip compressed data...
./websites/visual6502.org/wiki/skins/common/feed.css?270: gzip compressed data...
./websites/visual6502.org/wiki/index.php?title=Special:Contributions/EdS/index.html: gzip compressed data...
```

I decompressed these four files by hand. It looks like there are options you can give gunzip
to have it decompress into the same filename as the argument, but I manually renamed them .gz
files and then manually decompressed them.

2023/10/31

Yesterday I reorganized the code for mkmd and did a few other things.
Today I renamed the three files that contained single quotes
by replacing the single quotes with commercial @-signs.
I then made a list of all the HTML files in downloads:
```
find downloads -type f  | xargs file | grep 'HTML document' | sed 's/:  *HTML doc.*$//' > ALL_HTML_FILES
```
This found 174 HTML files, many of which are named unexpectedly (.jpg, .png, etc.)

I then put mkmd in my ~/bin as "md" and ran:
```
cat ALL_HTML_FILES | xargs -n 1 ~/bin/md -d  > /dev/null 2>foo
```
which produced 174 copies of `mkmd: success` in `foo`.
So the Golang HTML parser successfully parses all downloaded HTML files,
and the tree walker successfully prints them, or thinks it did.










