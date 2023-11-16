#!/bin/bash
#
# NAME
#   bld_auth.sh - build the authorship information in ./wiki/rdf
#
# USAGE:
#   ./tools/bld_auth.sh
#
# DESCRIPTION
#   This tool constructs the authorship information in wiki/rdf
#   from the RDF content in work/wiki/rdf. The content in work/wiki/rdf
#   was downloaded from the Wayback Machine using dl_auth.sh.

PROGNAME="./tools/bld_auth.sh"

SRCROOT="./work/wiki/rdf"
DSTROOT="./wiki/rdf"

echo "Building authorship information in wiki/rdf..."

rm -rf $DSTROOT ; mkdir $DSTROOT
if [ ! -d $DSTROOT ] ; then
	echo "failed to create $DSTROOT"
	exit 2
fi
cat $SRCROOT/README.md.prototype | grep -v 'this line is replaced' > $DSTROOT/README.md

ls $SRCROOT | while read rdffile ; do
	srcname=$SRCROOT/${rdffile}
	dstname=$DSTROOT/${rdffile}.md

	# first the file...
	(echo '```' ; echo)  > $dstname
	cat $srcname	    >> $dstname
	(echo '```' ; echo) >> $dstname

	# ...and then the link to it. We need to escape underscores
	# in the visible [ ] part of the text of the link.
	displayname=$( echo "${rdffile}" | sed 's/_/\\_/g' )
	echo "[${displayname}](./${rdffile}) " >> $DSTROOT/README.md
done
echo "Done"
exit 0
