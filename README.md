# golang-hex-dumper

A simple hex dump library routine for diagnostics that need delve into the binary. Displays bytes in hex.

Comes with example command line tool in the gohexdump directory.

`go get github.com/glycerine/golang-hex-dumper/...`

command line use:

`gohexdump <file to dump as hex>`

example:

~~~
$ gohexdump /tmp/binarystuff
pos 00  hex:  85  a1  44  92  a5  68  65  6c    '..D..hel'
pos 08  hex:  6c  6f  a5  77  6f  72  6c  64    'lo.world'
pos 16  hex:  a1  45  92  20  11  a1  47  91    '.E. ..G.'
pos 24  hex:  cb  c0  25  00  00  00  00  00    '..%.....'
~~~

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

