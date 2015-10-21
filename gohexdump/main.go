/*
// gohexdump: a simple hexdumper utility written in Golang
//
// Copyright 2015 Jason E. Aten <j.e.aten -a-t- g-m-a-i-l dot c-o-m>
// License: Apache 2.0
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	hex "github.com/glycerine/golang-hex-dumper"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gohexdump <binary file to display>\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	fn := os.Args[1]
	if !FileExists(fn) {
		fmt.Fprintf(os.Stderr, "hexdump error: file '%s' does not exist.\n", fn)
		usage()
	}

	by, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hexdump error: problem reading file '%s': '%s'\n", fn, err)
		os.Exit(1)
	}

	hex.Dump(by)
}

func FileExists(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}
