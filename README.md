# golang-hex-dumper

A simple hex dump library routine for diagnostics that need binary.

Comes with example command line tool in the gohexdump directory.

`go get github.com/glycerine/golang-hex-dumper/...`

command line use:

`gohexdump <file to dump as hex>`

library use:

~~~

import (
	hex "github.com/glycerine/golang-hex-dumper"
)
...

by := getSomeBytesYouNeedToInspect()
hex.Dump(by)

~~~

Copyright 2015 Jason E. Aten, Ph.D.
License: MIT
