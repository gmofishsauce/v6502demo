#!/bin/sh

# Iterate the files in RAW calling fix_one.sh on each, passing filename and target directory
# Usage: ./fixer.sh
# Run from root of repo
# Overwrites "fixed" files in $CONTENT (WaybackMachine header and footer removed)

# Subdirectory for raw downloads (input directory - don't touch)
RAW="d_1_raw"

# Subdirectory for pages with Wayback Machine header and footer, style sheets, scripts, etc. removed
FILTERED="d_2_filtered"

for i in $RAW/* ; do
	./tools/fix_one.sh $i $FILTERED
done
