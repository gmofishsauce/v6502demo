#!/bin/bash
#
# NAME
#   dl.sh - download the files in $1 from the Wayback Machine, slowly.
#
# USAGE:
#   ./tools/dl.sh file_list.jsonlines
#
# DESCRIPTION
#	This tool reads lines of JSON in a rigid format. It parses them clumsily,
#   using sed, to obtain the file_url. It then attempts to download the file
#   using the wayback machine downloader (Ruby gem; search Github). See
#   jsonlines.org for a description of the JSON Lines format.
#
#   The reason we don't use the Wayback Machine Downloader to grab the whole
#   site is that WMD doesn't provide control over delay between requests.
#   The WM soon starts refusing to connect. Note that it didn't do this the
#   very first time I downloaded the entire Visual6502.org wiki, but it did
#   every time after that; I think they're sophisticated about throttling.
#
#   The JSON lines file is expected to be in the format produced by Wayback
#   Machine and emitted by the wayback_machine_downloader Ruby gem with the
#   --list option. An example of a suitable line, broken here for clarify:
#     {"file_url":"http://visual6502.org/wiki/skins/common/history.js?270",
#      "timestamp":"20210519010639","file_id":"wiki/skins/common/history.js?270"},
#
#   The command used to generate the JSON for the actual download:
#     wayback_machine_downloader http://visual6502.org --only wiki \
#     --exclude "/\&[A-Za-z]+|Special:/" --list > ALL_FILES.json
#   After which the opening [ and closing ] were removed manually.
#
# BUGS
#   This script parses JSON using sed. Get past it.
#   This script is probably Macos-specific (not Gnu sed).

PROGNAME="./tools/dl.sh"
if [ "x$1" = x ] ; then
	echo "Usage: $PROGNAME file_list.sh"
	exit 1
fi

if [ ! -r $1 ] ; then
	echo "$PROGNAME: $1: cannot open"
	exit 2
fi

INFILE=$1

# Because we use the WMD to do the download, instead of using e.g. wget,
# we don't need to know the actual Wayback Machine URL of the download.
# Some of the tricks to construct the actual URL are contained in the
# Golang code found in ./tools, though.

TRUE=0
RESULTFILE=wmd_stdout.txt

while TRUE ; do
	url=$( head -1 "$INFILE" | sed -e 's/^.*"file_url":"//' -e 's/",".*$//' )
	if [ "x$url" = "x" ] ; then
		echo "DONE"
		break
	fi

	backoff=$( jot -r -n 1 5 15 )
	todo=yes

	while [ "$todo" = "yes" ] ; do
		wayback_machine_downloader -e "${url}" > $RESULTFILE

		# Typical error:
		# websites/visual6502.org/wiki/skins/common/IE80Fixes.css?270 was empty and was removed.

		if grep 'was empty and was removed' $RESULTFILE ; then
			todo=yes
			((backoff*=2))
			echo "error: backoff doubled"
		else
			todo=no
			lines=$( wc -l $INFILE | awk '{print $1}' )
			echo $lines lines remaining
			lines=$( expr $lines - 1 )
			tail "-$lines" $INFILE > $INFILE.tmp
			mv $INFILE.tmp $INFILE
		fi
		echo sleep $backoff
		sleep $backoff
	done
done
