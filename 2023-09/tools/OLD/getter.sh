#!/bin/sh

# Usage: cat ALL_PAGES | ./tools/getter.sh
# The ALL_PAGES file was made by hand from the Special All Pages HTML page
# Run this from root of repo
# Places downloaded files in $RAW
# Then run the Fixer to do the first step in processing toward .md files.

# Subdirectory for raw downloads
RAW="d_1_raw"

# String to get the particular version of each URL from the Wayback Machine
URLBASE="https://web.archive.org/web/20210405071236"

while read url ; do
	name=$( echo $url | sed 's#^.*/##' | sed 's/%../_/g' | tr '.?=&,-)(' _ )
	name="${name}.html"

	if [ -r $RAW/$name ] ; then
		echo "$name: already downloaded."
		continue
	fi

	url="${URLBASE}/$url"
	echo "Fetching $url to $name ..."

	wget $url -O $name
	mv $name $RAW

	if ! file $RAW/$name | grep 'HTML document text' ; then
		echo "-----------------------------------------"
		echo "WARNING: $RAW/$name: not an HTML document"
		echo "-----------------------------------------"
	fi

	echo '-----------'
	echo "Pausing ..."
	echo '-----------'
	sleep $(jot -r 1 6 12)
done
