## Welcome to the v6502 Wiki Recovery Repo

The recovered Visual6502.org wiki is being built in [`wiki`](./wiki).
The "recovered wiki" will not be a wiki but rather a static, markdown-base website
that will be rendered to HTML with Jekyll. The initial, demo-only deployment will
be on Github Pages.

The original Wiki states:

Content is available under [Attribution-NonCommercial-ShareAlike 3.0 Unported](https://web.archive.org/web/20210405071423/http://creativecommons.org/licenses/by-nc-sa/3.0/) license.

To honor this license, I must list the authors.
This information is found in [`./wiki/rdf`](wiki/rdf).

## Process

- The original content is downloaded from the Wayback Machine
- The files are renamed to avoid illegal characters in URLs.
- A markdown-based equivalent of the entire wiki is generated
- Github Pages renders the markdown content as a static HTML site

Each of these steps is described in more detail below.

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

### Download the authorship information

The license on the wiki requires listing the authors. I found that all authorship information
was present in the Wayback Machine as a set of XML (RDF) files, one per content file. The RDF
files are referenced by `<link>` tags in each HTML document.

Code in the mkmd tool can be used to identify and download the most recent version of each RDF
file in the Wiki. These files were downloaded to [`./wiki/rdf`](wiki/rdf). This is done a single
file at a time, using multiple invocations of the command the command

```
mkmd -r -o wiki/rdf html_file.ext
```

I used `find` to locate all the HTML documents in the base wiki and a one-line shell loop
to run mkmd per file. I found 174 HTML files, many of which have file extensions like `.jpg`
or `.png`. From this list I was able to download 161 RDF files.

I build `mkmd` in the `./tools` directory like this: `go build -o mkmd`.

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

