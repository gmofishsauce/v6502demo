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
