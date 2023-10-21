#!/bin/sh

# Usage: get_img_for_file.sh
# Check the images linked from an HTML file given by the first argument
# If they don't exist in the directory named by the second argument,
# get them using wget and store them in that directory.
#
# Note: this program searches for links to certain file extensions. This
# program doesn't know anything about the content of the downloaded files.
# In the visual6502 Wiki, some files with the extension .jpg are actually
# HTML files. This is dealt with by other tools.

# String to get the particular version of each URL from the Wayback Machine
URLBASE="https://web.archive.org"

if test $# != 2 ; then
	echo "Usage: get_img_for_file dir/infile outdir"
	exit 1
fi

if ! test -r ${1} ; then
	echo "Cannot open: ${1}"
	exit 1
fi
infilepath=${1}

if ! test -d ${2} ; then
	echo "Not a directory: ${2}"
	exit 1
fi
targetdir=${2}

# Find links to files that have image file names (as noted, some of these turn
# out to be HTML files which we handle in a later processing step). Split the
# file on spaces into many lines because this simple approach can't handle two
# or more links in the same line.
tmpfile=get_img_for_file_$$.tmp
trap "{ rm -f $tmpfile ; exit 3 ; }" SIGINT SIGQUIT SIGTERM
trap "{ rm -f $tmpfile ; exit 0 ; }" EXIT

<$infilepath tr ' ' '\n' |\
             grep -e '.*.png' -e '.*.jpg' -e '.*.gif' |\
             sed -e 's/^.*src="//' -e 's/^.*href="//' -e 's/".*$//' |\
			 grep -v 'alt=' >> ${tmpfile}

for i in $(cat ${tmpfile}) ; do
	if echo $i | grep 'https://web.archive.org' ; then
		echo "BARE URL USED"
		url=$i
	else
		echo "ROOT URL ADDED"
		url="${URLBASE}/$i"
	fi
	name=$(basename $i)
	target=${targetdir}/${name}

	if test -r ${target} ; then
		echo "File exists: ${target}"
		sleep 1
	else
		wget ${url} -O ${target}
		echo "Pausing ..."
		sleep $(jot -r 1 4 12)
	fi
done

exit 0
