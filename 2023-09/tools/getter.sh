#!/bin/sh

# Usage: cat ALL_PAGES | ./tools/getter.sh
# Run from root of repo
# Places downloaded files in $RAW
# Places reduced files in $CONTENT (WaybackMachine header and footer removed)

# Subdirectory for raw downloads
RAW="./raw"

# Subdirectory for pages with Wayback Machine header and footer, style sheets, scripts, etc. removed
CONTENT="./content"

# DOCTYPE string from typical page
DOCTYPE='<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">'

# String to get the particular version of each URL from the Wayback Machine
URLBASE="https://web.archive.org/web/20210405071236"

while read url ; do
	name=$(echo $url | sed 's#^.*/##' | sed 's/%../_/g' | tr '.?=&' _)
	name="${name}.html"

	url="${URLBASE}/$url"
	echo "Fetching $url to $name ..."

	wget $url -O $name
	mv $name $RAW

	echo $DOCTYPE                      > $CONTENT/$name
	echo '<head>'                     >> $CONTENT/$name
	echo '<html lang="en" dir="ltr">' >> $CONTENT/$name
	echo '<head>'                     >> $CONTENT/$name
	grep '<title>' $RAW/$name         >> $CONTENT/$name
	echo '</head>'                    >> $CONTENT/$name
	echo '<body>'                     >> $CONTENT/$name

	<$RAW/$name sed -n '/-- start content --/,/-- end content --/p' >> $CONTENT/$name

	echo '</body>'                    >> $CONTENT/$name
	echo '</html>'                    >> $CONTENT/$name

	DATE=$(date)
	echo "<!-- written by getter $DATE -->" >> $CONTENT/$name

	echo "Pausing ..."
	sleep $(jot -r 1 2 6)
done
