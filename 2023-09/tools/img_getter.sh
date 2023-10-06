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

# First pass - get all the img src= links into ALL_IMG_LINKS
cp /dev/null $ALL_IMG_LINKS
while read filename ; do
	<$filename grep 'src=' |\
	sed -e 's/^.*src="//' -e 's/".*$//' |\
	sort | uniq >> $ALL_IMG_LINKS
done

for i in $(cat $ALL_IMG_LINKS) ; do
	url="${URLBASE}/$i"
	name=$(basename $i)

	echo "Fetching $url to $name ..."
	wget $url -O $RAWIMG/$name

	echo "Pausing ..."
	sleep $(jot -r 1 4 12)
done
