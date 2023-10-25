package main

import (
	"fmt"
	"os"
	"runtime"
)

// Print debug output. This is just fprintf to stderr.
func dbg(s string, args... any) {
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
        fmt.Printf("TODO called from %s\n", details.Name())
		todoDone[details.Name()] = true
    }
	return nil
}

