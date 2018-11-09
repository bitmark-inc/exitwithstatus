// Copyright (c) 2014-2018 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// replacement for os.Exit(N) that runs defers
//
// Original article: http://soniacodes.wordpress.com/2011/04/28/deferred-functions-and-an-exit-code/
//
// This is modified version to pass an integer status to os.Exit
//
// usage: as the first line of main add:
//   defer exitwithstatus.HandleFatal()
//
// to exit with a particular integer value:
//   exitwithstatus.Exit(42)
package exitwithstatus

import (
	"fmt"
	"os"
	"path/filepath"
)

type fatal struct {
	err interface{}
}

// Exit with the integer status value
//
// e.g to exit calling all defers with the value three: exitwithstatus.Exit(3)
func Exit(status int) {
	panic(fatal{status})
}

// This must be the first defer in the program
//
// i.e. the first line of main must be:  defer exitwithstatus.Handle()
func Handler() {
	if err := recover(); err != nil {
		if status, ok := err.(fatal); ok {
			s := status.err.(int)
			os.Exit(s)
		}
		panic(err) // this is some other kind of panic so pass it up the chain
	}
}

// print a message to stderr and exit wit error status
func Message(message string, args ...interface{}) {

	m := "failed"

	// if a blank message print a simple usage message
	if message == "" {

		// tidy up the program name
		programName := filepath.Base(os.Args[0])

		m = fmt.Sprintf("Use %s -h to show usage", programName)
	} else {

		// user supplied message
		m = fmt.Sprintf(message, args...)
	}

	// print to stderr with final '\n'
	fmt.Fprintln(os.Stderr, m)

	Exit(1)
}
