# Welcome to the visual6502.org Wiki Recovery Repo

Between about 2007 and 2015, a team of dedicated hobbyists did significant
reverse engineering work on several early microprocessors, particularly the
venerable 6502. Their work can be found at [visual6502.org](http://visual6502.org).

The team maintained a wiki that used MediaWiki technology. The wiki is no longer
functional [as you can see](http://visual6502.org/wiki).

This repository represents my personal effort to restore this content to easy
availability on the internet.

This site, however, is historical documentation; it is not a wiki. The recovered
site exists as a [static, markdown-base website](https://gmofishsauce.github.io/v6502demo/)
in Github Pages. The process used to build the site is described below.

The original Wiki states:

Content is available under [Attribution-NonCommercial-ShareAlike 3.0 Unported](https://web.archive.org/web/20210405071423/http://creativecommons.org/licenses/by-nc-sa/3.0/) license.

To honor this license, I must list the authors.
This information is found in [`./wiki/rdf`](wiki/rdf).
Note that the links at the bottom of the README file for that directory
are functional only in the rendred Github Pages site, not on Github proper.

## Process

- The original content was downloaded from the Wayback Machine.

This copy never modified. It is located in [websites](./websites).

- Additions are made and various cleanup changes are made.

This copy is located in [work](./work).

- A static markdown representation is built. It is located in [wiki](./wiki).

This copy is located in [wiki](./wiki).

- Github Pages uses Jekull to render the markdown content as a static HTML site.

The rendered site is at [https://gmofishsauce.github.io/v6502demo/](https://gmofishsauce.github.io/v6502demo/).

## Tooling

The Wayback Machine Downloader (WMD) was used to download the Wiki.

Additional processing was done a custom tool written in Golang, `mkmd`.
The source code for this tool is found in the [`tools`](./tools) subdirectory.
I am the sole author of the mkmd tool. It is GPLv3 licensed.

## Recovery Details

### Download the base wiki using the Wayback Machine Downloader

The WMD is a [well-documented Ruby Gem](https://github.com/hartator/wayback-machine-downloader).
Because the Wiki contains many internal links and the Wayback Machine blindly chases all of them,
the WMD by default downloads tens of thousands of pages. Specific exclusions reduced this to less
than 1000 pages (including all images) for the entire V6502 Wiki:
```
wayback_machine_downloader http://visual6502.org --only wiki \
  --exclude "/\&[A-Za-z]+|Special:/" > all_files.wmd 2>stderr.wmd
```

The command above can be checked without doing an actual download by adding the command line
option `--list` to the wayback_machine_downloader command line. The list of files to be
downloads is written to the standard output (e.g. `all_files.wmd` in the example above).

### (No, really) download the files using the Wayback Machine Downloader

After my initial download, I found that I could never download the entire site again (which
turned out to be required because I did not treat the download as pristine). It seemed like
a throttling issues. The Wayback Machine Downloader does not have any options for adding
delays between the individual file downloads.

In response I wrote a shell script, `./tools/dl.sh`. To use it, run the WMD command given
above with the `--list` option. This produces a JSON file which I committed as ALL_FILES.json.
Make a copy of the file (e.g. `lines.json`) and remove the first line '[' and the last line ']'
from the file; this leaves a file n "JSON lines" format (see jsonlines.org). Run
`./tools/dl.sh lines.json` to download all the files. The script will incrementally remove lines
from the file as it successfully downloads each file; this allows the script to be killed and
restarted if issues occur. The script delays 5 to 15 seconds between each pair of file downloads.
I found this to be sufficient and did not investigate further.

The WMD downloads into `websites/hostname`, `websites/visual6502.org` in this case. But I
found that this second download pulled two variations on the name "visual6502.org", each
containing a couple of files. These can be seen in `websites/`. These fake names complicate
processing. I decided to defer dealing with these. I created a directory `work/` and copied
`./websites/visual6502.org` recursively to `work`.  From here until the cleanup at the end
of this process, "the files" (or "all the files") refers to the content of `./work` only.
The `websites` directory is kept pristine and the two side directories are addressed at the
end of the process.

### A note about file naming

All processing in this part of the "pipeline" must be done with extreme care because of the
file naming. The downloaded files were named from their wiki page titles, which can include
any character. Three of the pages in this wiki contain single quotes in their titles which
become single quotes in their downloader filenames, and a few other files contain other shell
metacharacters in their names. Fixing these is not trivial, because links to these files
contain URL-encoded representations of these metacharacters, which must also be fixed. This
is done later and details are given there.

### Decompress the working files

Some non-image files in the wiki are gzipped. Image files are sensibly not gzipped as they
don't compress much. I wrote another shell script, `./tools/gunz.sh`, to find all the gzipped
files and unzip them. The script is idempotent (it can be safely rerun if gzipped files are
ever added to the repo).  The script uses tricks like `find -print0 ... | xargs -0 ...` because
of the file naming issue, which is resolved in a later step.

### Download the authorship information

The license on the wiki requires listing the authors. I found that all authorship information
was present in the Wayback Machine as a set of XML (RDF) files, one per content file. The RDF
files are referenced by `<link>` tags in each HTML document.

Code in the mkmd tool can be used to identify and download the most recent version of each RDF
file in the Wiki. These files were downloaded to [`./wiki/rdf`](wiki/rdf). This is done a single
file at a time, using multiple invocations of the command

```
mkmd -r -o wiki/rdf html_file.ext
```

I wrote `tools/dl_auth.sh` to download all authorship information into the [work](./work)
directory. This script runs `mkmd -r` for each file. Note that many of the HTML files have
file extensions `.png` or `.jpg`.

I build `mkmd` in the `./tools` directory like this: `go build -o mkmd`. The actual URL of
each of these `.rdf` files within the Wayback Machine is peculiar; the Go code knows the rule
for constructing it.

### Rename all the files

Every downloaded file name will end up in the URL of a file rendered into the Gibhub Pages site
by the Jekyll processor. Due to the structure of the wiki, essentially all the files (including
some of the RDF files) have names that cannot be used in URLs. Many of them are handled correctly
in practice even though they form technically illegal URLs; others causes problems when the Github
Pages pipeline runs Jekyll to render the markdown I generate as static HTML.

I renamed all the offending files (i.e. most of the files) according to standard rules. This is
done using another command line option to mkmd (description TBD). The renaming of course breaks
all the links to the renamed documents. All URLs are rewritten in later steps by applying the
same character remapping rules. Unlike the file names, the URLs that reference them are usually
URL encoded. The URLs must be URL decoded, then fixed, and then URL encode (although the last
step is likely unnecessary with all the special characters substituted for legal URL characters).

### Set up Github Pages and create some README files

I enabled Github Pages for the entire repo. Much of the repo is not accessible through the Pages
site, because it only renders markdown. I created some README.md files, which become index.html
files in effect.

### Move the images

The site will end up in `v6502demo/wiki`, so the images need to be there. Rather than copy the
ones I downloaded which would double the size of the repo, I moved the downloadeded `images/`
directory from `downloads/visual6502.org/wiki/images` to `wiki/images`.

### Write the .md files

This is the point of the whole exercise: writing a static markdown file that captures all the
valuable content of the original MediaWiki page. This step includes updating all the URLs that
may have pointed to renamed files. All the hrefs in the wiki fall into one of the catagories
described by the following sequence of `grep -v` commands:
```
grep href * |\
	grep -v 'href="/wiki' |\
	grep -v "api.php" |\
	grep -v 'href="#[A-Za-z]' |\
	grep -v 'http://visual6502' |\
	grep -v "http://[A-Za-z]" |\
	grep -v "https://[A-Za-z]" |\
	grep -v 'href="/favicon.ico"' |\
	grep -v 'href="#[0-9]'
```
The first line alone produces over 8500 lines of output, of which the `grep -v` commands
filter all but one line, an FTP link. The markdown generator will have to deal with all
the filtered URL types when they are encountered (they may not be, because the hrefs may
be found in HTML tags that the markdown processor ignores; note for example the "api.php"
links to the MediaWiki API).

