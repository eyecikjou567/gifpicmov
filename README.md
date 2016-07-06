# GifPicMover

This handy little tool will pic out any .gif files in a given search directory and put them in a single target directory.

This is especially useful when you want to run batch tools over your picture collection that have no .gif support.

## Building

Building the Binary was tested under Go 1.6.2

Dependencies:

* golang.org/x/crypto/sha3

Simply `go get` and then `go build` to get the gifpicmov binary.

## Usage

By default GPM will not execute actions, only log them. Check the output before you disable the safeguard.

Commandline Options:

* `-dry` - Default: true, Determines if GPM will do a dry run, only log but not execute actions
* `-keep-name` - Default: false, If set to true, GPM won't rename the file, if set to false, GPM will set the filename to the SHA3-256 Hash of the GIF
* `-tar` - Default: "./0gif", Directory in which gifs will be placed
* `-dir` - Default: ".", Directory in which gifs are searched, including all subfolders and filesystems