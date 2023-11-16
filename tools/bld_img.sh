#!/bin/bash
#
# NAME
#   bld_img.sh - build the image hierarchy in wiki/images
#
# USAGE:
#   ./tools/bld_img.sh
#
# DESCRIPTION
#   This tool constructs the image data hierarchy for the entire wiki in
#   ./wiki/images from the prototype in work/wiki/images using a single
#   call to rsync. The work/wiki/images hiearchy is authoritative; it
#   began as a copy of websites/visual6502.org/wiki/images, but has been
#   hand-maintained (manually enhanced with additional content) since the
#   original download.

PROGNAME="./tools/bld_img.sh"

# Note that rsync is sensitive to trailing slashes on paths.
SRCROOT="./work/wiki/images"
DSTROOT="./wiki/"

echo "Building images in wiki/images..."
rsync -avz --delete $SRCROOT $DSTROOT
exit $?
