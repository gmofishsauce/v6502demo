#!/bin/bash
#
# NAME
#   gunz.sh - unzip the files in $1
#
# USAGE:
#   ./tools/gunz.sh directory_containing_files
#
# DESCRIPTION
#   gunz finds all files in the directory that are actually gzipped. For each
#   gzipped file, it checks to see if the file ends with .gz and if not, it
#   renames the file. It then gunzips the file. Because it finds files that
#   appear (to the "file" command) to be gzipped files, it is idempotent (it
#   can be rerun if additional gzipped files are added to the directory). To
#   be clear, the files are unzipped "in place".
#
# BUGS
#   Yes, I know there are command line options to allow gunzipping a file with
#   an arbitrary extension. This seemed like the most straightforward answer.

PROGNAME="./tools/gunz.sh"
if [ "x$1" = x ] ; then
	echo "Usage: $PROGNAME directory"
	exit 1
fi

if [ ! -d "$1" ] ; then
	echo "$PROGNAME: $1: not a directory"
	exit 1
fi
SRCDIR="$1"

# The target directory for the downloads. Since the original site was a
# MediaWiki, we can't know anything about the actual layout of content in
# files, or even whether there were files, versus e.g. a database holding
# the content. The general approach I'm now taking for recovery is to
# download the original files into a directory websites/ which I then
# keep pristine. This download doesn't include the authorship information
# because the Wayback Machine Downloader doesn't find it. I then build up
# a second copy of the site in work/ which includes various processing
# including decompression, downloading of the authorship information, 
# renaming files to avoid illegal URLs, and hand fixes to clean up one
# of a kind data issues. Finally the work/ directory is processed into
# wiki/site as markdown plus images, which is then rendered back to a
# static HTML site by the Jekyll processor in Github Pages. As with many
# of the tools used to do all this, the TARGET assumes you are running
# from the root of the repo.

ALL_FILES=$( find "$SRCDIR"  -print0 | xargs -0 file |\
			grep 'gzip compressed data' | sed 's/: .*$//' )
for filename in $ALL_FILES ; do
	echo "decompress $filename..."
	mv "${filename}" "${filename}.gz"
	if gunzip "${filename}.gz" ; then
		sleep 3
	else
		echo "gunzip failed"
		exit 2
	fi
done
exit 0
