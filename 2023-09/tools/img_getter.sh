#!/bin/sh

# Usage: ls content/* | ./tools/img_getter.sh
# Run from root of repo
# Places a list of image links in ALL_IMG_LINKS
# Places downloaded files in $RAWIMG

# Subdirectory for raw downloads
RAWIMG="./rawimg"

# String to get the particular version of each URL from the Wayback Machine
URLBASE="https://web.archive.org"

# File containing all the image links
ALL_IMG_LINKS="./ALL_IMG_LINKS"

# File containing names of files that end with .jpg but are actually HTML
ALL_HTML_FILES_NAMED_JPG="./ALL_HTML_FILES_NAMED_JPG"

# First pass - get all the img src= links into ALL_IMG_LINKS
cp /dev/null $ALL_IMG_LINKS
cp /dev/null $ALL_HTML_FILES_NAMED_JPG

while read filename ; do
	<$filename grep -e '.*.png' -e '.*.jpg' |\
	sed -e 's/^.*src="//' -e 's/^.*href="//' -e 's/".*$//' |\
	sort | uniq >> $ALL_IMG_LINKS
done

for i in $(cat $ALL_IMG_LINKS) ; do
	url="${URLBASE}/$i"
	name=$(basename $i)

	echo "Fetching $url to $name ..."
	wget $url -O $RAWIMG/$name

	if file $RAWIMG/$name | grep 'HTML document text' ; then
		echo "HTML file: $RAWIMG/$name"
		echo $RAWIMG/$name >> $ALL_HTML_FILES_NAMED_JPG
	fi

	echo "Pausing ..."
	sleep $(jot -r 1 4 12)
done
