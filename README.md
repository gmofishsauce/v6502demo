## Welcome to the v6502 Wiki Recovery Repo

The recovered Visual6502.org wiki is being built in [`wiki`](./wiki).
The "recovered wiki" will not be a wiki but rather a static, markdown-base website
that will be rendered to HTML with Jekyll. The initial, demo-only deployment will
be on Github Pages.

The original Wiki states:

Content is available under [Attribution-NonCommercial-ShareAlike 3.0 Unported](https://web.archive.org/web/20210405071423/http://creativecommons.org/licenses/by-nc-sa/3.0/) license.

To honor this license, I must list the authors.
This information is found in [`./wiki/rdf`](wiki/rdf).

## Tooling

The Wayback Machine Downloader (WMD) was used to download the Wiki.

Additional processing was done a custom tool written in Golang, `mkmd`.
The source code for this tool is found in the [`tools`](./tools) subdirectory.
I am the sole author of the mkmd tool. It is GPLv3 licensed.

## High level recovery steps

### Download using the Wayback Machine Downloader

The WMD is a [well-documented Ruby Gem](https://github.com/hartator/wayback-machine-downloader).
Because the Wiki contains many internal links and the Wayback Machine blindly chases all of them,
the WMD by default downloads tens of thousands of pages. Specific exclusions reduced this to less
than 1000 pages (including all images) for the entire V6502 Wiki:
```
wayback_machine_downloader http://visual6502.org --only wiki \
  --exclude "/\&[A-Za-z]+|Special:/" > all_files.wmd 2>stderr.wmd
```

### Obtain the authorship information

The mkmd tool was used to identify and download the most recent version
of each Resource Description (`.rdf`) file in the Wiki.
These files are found in [`./wiki/rdf`](wiki/rdf). In Chrome, at least,
the `.rdf` pages download, rather than displaying in a tab.

### Set up Github Pages and create some README files

I reenabled GH Pages for the entire repo. Much of the repo is not accessible through the Pages
site, because it only renders markdown. I created some README.md files, which become index.html
files in effect; these READMEs may be temporary.

### Move the images

The site will end up in `v6502demo/wiki`, so the images need to be there. Rather than copy the
ones I downloaded which would double the size of the repo, I moved the downloadeded `images/`
directory from `downloads/visual6502.org/wiki/images` to `wiki/images`.

