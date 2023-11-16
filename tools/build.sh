#!/bin/bash
#
# NAME
#   build.sh - build the static markdown site in ./wiki from ./work
#
# USAGE:
#   ./tools/build.sh
#
# DESCRIPTION
#   This tool processes the files in work/ to create the markdown site
#   in wiki/. All the file paths are hardwired. The files in work/ must
#   have been carefully processed to remove shell metacharacters before
#   running this script; see the README file at the top of the repo for
#   details.

PROGNAME="./tools/build.sh"

SRCROOT="./work/wiki"
DSTROOT="./wiki"

# Every build builds clean
rm -rf ${DSTROOT}/*

# Build the top-level README
cp $SRCROOT/README.md.prototype $DSTROOT/README.md

# Build the authorship information in wiki/rdf
echo "Building authorship information in wiki/rdf..."
DSTDIR=${DSTROOT}/rdf
SRCDIR=${SRCROOT}/rdf

mkdir $DSTDIR
if [ ! -d $DSTDIR ] ; then
	echo "failed to create $DSTDIR"
	exit 2
fi
cat $SRCDIR/README.md.prototype | grep -v 'this line is replaced' > $DSTDIR/README.md

ls $SRCDIR | while read rdffile ; do
	# first the file...
	srcname=$SRCDIR/${rdffile}
	dstname=$DSTDIR/${rdffile}.md
	cat $srcname	  >> $dstname
	# ...and then the link to it. We need to escape underscores
	# in the visible [ ] part of the text of the link.
	displayname=$( echo "${rdffile}" | sed 's/_/\\_/g' )
	echo "[${displayname} ](./${rdffile})" >> $DSTDIR/README.md
done

echo "done"
echo "Building source files in wiki..."
echo "TODO"
echo "done"

exit 0
