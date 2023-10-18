#!/bin/bash

# Hoist image files that are really HTML files into a directory.
TARGETDIR=d_4_hoisted

for i in d_3_img/* ; do
	if file $i | grep 'HTML document text' > /dev/null ; then
		echo "Hoisting $i"
		./tools/fix_one.sh $i ${TARGETDIR}
		name=$(basename $i)
		mv ${TARGETDIR}/$name ${TARGETDIR}/${name}.html
		sleep 3
	else
		echo "Not an HTML document: $i"
		sleep 1
	fi
done
