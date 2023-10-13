#!/bin/sh

# Fix one html file by changing HTML entities to XML and removing stuff added by Wayback Machine
# Usage: fix_one.sh inputPathName outputDirectory
# Normally, inputPathName must be dir/file

# DOCTYPE string from typical page
DOCTYPE='<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">'

if test $# != 2 ; then
	echo "Usage: fix_one dir/infile outdir"
	exit 1
fi

if ! test -r ${1} ; then
	echo "Cannot open: ${1}"
	exit 1
fi
infilepath=${1}
filename=$(basename ${1})

if ! test -d ${2} ; then
	echo "Not a directory: ${2}"
	exit 1
fi
outputfile=${2}/$filename

echo $DOCTYPE                      > $outputfile
echo '<html lang="en">'           >> $outputfile
echo '<head>'                     >> $outputfile
grep '<title>' $infilepath        >> $outputfile
echo '</head>'                    >> $outputfile
echo '<body>'                     >> $outputfile

# output the body content and translate HTML entities to XML
<$infilepath sed -n '/-- start content --/,/-- end content --/p' |\
			 sed -e "s/\&nbsp;/\&#160;/g"   >> $outputfile

echo '</body>'                    >> $outputfile
echo '</html>'                    >> $outputfile

DATE=$(date)
echo										>> $outputfile
echo "<!-- written by fix_one.sh $DATE -->" >> $outputfile
