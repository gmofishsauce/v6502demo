#!/bin/bash
#
# NAME
#   dl_auth.sh - download the authorship (RDF) files for the wiki
#
# USAGE:
#   ./tools/dl_auth.sh directory_containing_HTML_files target_directory
#
# DESCRIPTION
#   dl_auth finds the HTML files in the argument directory. It
#	then runs mkmd, which searches the file for a <link> tag with
#   specific attributes. If such a tag is found, mkmd tries to
#   construct a Wayback Machine URL for the RDF file mentioned
#   in the link and download it.
#
# BUGS
#   Probably.

PROGNAME="./tools/dl_auth.sh"
if [ "x$1" = x -o "x$2" = x ] ; then
	echo "Usage: $PROGNAME HTML_directory RDF_target_directory"
	exit 1
fi

if [ ! -d "$1" -o ! -d "$2" ] ; then
	echo "$PROGNAME: both arguments must be directories"
	exit 1
fi
SRCDIR="$1"
DSTDIR="$2"

# At this point in the "processing pipeline" from MediaWiki to
# static website, downloaded filenames have not yet been fixed
# and may so contain shell metacharacters, even unbalanced quotes.
# Fortunately, none of them appear to contain double quotes.

ALL_FILES=$( find "$SRCDIR" -print0 | xargs -0 file |\
			grep HTML | sed 's/: .*$//' )

for filename in $ALL_FILES ; do
	echo ./tools/mkmd -r -o "$DSTDIR" "$filename"
	./tools/mkmd -r -o "$DSTDIR" "$filename"
	backoff=$( jot -r -n 1 5 15 )
	sleep $backoff
done
exit 0
