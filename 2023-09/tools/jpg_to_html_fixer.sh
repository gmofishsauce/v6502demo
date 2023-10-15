#!/bin/sh

# img_getter.sh - get the images for all (HTML) files in a source directory
#
# ./tools/img_getter.sh srcdir targetdir
#
# This tools processes the files in srcdir on the assumption they are HTML.
# It identifies all the image links (to .png, .jpg, or .gif) files and checks
# for the existence of a name-matching file in the target directory. If no
# file matching the name is in the target directory, the image is pulled with
# wget. Note that the image file may not actually be an image; in the Wiki,
# some aren't. This tool doesn't deal with that at all, it just pulls files
# with extensions into the target directory.

# String to get the particular version of each URL from the Wayback Machine
URLBASE="https://web.archive.org"

# First pass - get all the img src= links into ALL_IMG_LINKS
cp /dev/null $ALL_IMG_LINKS
cp /dev/null $ALL_HTML_FILES_NAMED_JPG

while read filename ; do
	<$filename grep -e '.*.png' -e '.*.jpg' |\
	sed -e 's/^.*src="//' -e 's/^.*href="//' -e 's/".*$//' >> ${ALL_IMG_LINKS}.tmp
done
sort ${ALL_IMG_LINKS}.tmp | uniq >> $ALL_IMG_LINKS
rm -f ${ALL_IMG_LINKS}.tmp

exit 5

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
