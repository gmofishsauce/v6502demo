package main

/*
Author: Jeff Berkowitz
Copyright (C) 2023 Jeff Berkowitz

This file is part of mkmd.

mkmd is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation, either version 3
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see http://www.gnu.org/licenses/.
*/

import (
	"fmt"
	"os"
	"runtime"
)

// Print debug output. This is just fprintf to stderr.
func dbg(s string, args... any) {
    pc, _, _, ok := runtime.Caller(1)
    details := runtime.FuncForPC(pc)
	where := "???"
    if ok && details != nil {
		where = details.Name()
    }
	s = "[at " + where + "]: " + s + "\n"
	fmt.Fprintf(os.Stderr, s, args...)
}

// Print debug output. This is just fprintf to stderr.
func dbg2(s string, args... any) {
    pc, _, _, ok := runtime.Caller(2)
    details := runtime.FuncForPC(pc)
	where := "???"
    if ok && details != nil {
		where = details.Name()
    }
	s = "[at " + where + "]: " + s + "\n"
	fmt.Fprintf(os.Stderr, s, args...)
}

var todoDone = make(map[string]bool)

// This function prints the callers name and TODO once per
// execution of the calling program. Arguments are ignored
// and are provided to make reference to unreference variables
// in a partially completely implementation.
func TODO(args... any) error {
	pc, _, _, ok := runtime.Caller(1)
    details := runtime.FuncForPC(pc)
    if ok && details != nil && !todoDone[details.Name()] {
        dbg("TODO called from %s", details.Name())
		todoDone[details.Name()] = true
    }
	return nil
}

