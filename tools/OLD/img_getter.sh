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

if test $# != 2 ; then
	echo "Usage: img_getter srcHtmlDir targetImgDir"
	exit 1
fi

if ! test -d ${1} ; then
	echo "Not a directory: ${1}"
	exit 1
fi
srcdir=${1}

if ! test -d ${2} ; then
	echo "Not a directory: ${2}"
	exit 1
fi
targetdir=${2}

for i in $(ls $srcdir) ; do
	echo "Processing ${srcdir}/$i"
	tools/get_img_for_file.sh "${srcdir}/$i" "${targetdir}"
	if [ $? = 3 ] ; then
		echo "Aborted."
		exit 2
	fi
done

echo "Done"
exit 0
