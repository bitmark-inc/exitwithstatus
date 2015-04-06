// Copyright (c) 2014-2015 Bitmark Inc.
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
	"strings"
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

// print a usage message and exit
func Usage(message string, args ...interface{}) {
	// display the version and exit if the version flag was specified.
	programName := filepath.Base(os.Args[0])
	programName = strings.TrimSuffix(programName, filepath.Ext(programName))
	usageMessage := fmt.Sprintf("Use %s -h to show usage", programName)

	if message != "" {
		fmt.Fprintf(os.Stderr, message, args...)
	}
	fmt.Fprintln(os.Stderr, usageMessage)
	Exit(1)
}
