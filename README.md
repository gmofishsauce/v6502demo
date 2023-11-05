## Welcome to the v6502 Wiki Recovery Repo

The recovered Visual6502.org Wiki is being built in [`mdsite`](./mdsite).
The "recovered" Wiki will not be a wiki but rather a static, markdown-base website
that will be rendered to HTML with Jekyll. The initial, demo-only deployment will
be on Github Pages.

The original Wiki states:

Content is available under [Attribution-NonCommercial-ShareAlike 3.0 Unported](https://web.archive.org/web/20210405071423/http://creativecommons.org/licenses/by-nc-sa/3.0/) license.

To honor this license, I must list the authors.
This information is found in `./mdsite/rdf`.

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
Yes, the log files are WMDs ;-)

### Obtain the authorship information

The mkmd tool was used to identify and download the most recent version
of each Resource Description (".rdf") file in the Wiki.
These files are found in [`./mdsite/rdf`](./mdsite/rdf).

