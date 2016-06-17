/*
eyeball is a simple library that watches a Go project and automatically builds and runs it when it detects that a file has been saved.

It is a fork of https://github/pilu/fresh, made to be much simpler.
*/
package main

import (
	"github.com/kevwu/eyeball/runner"
)

func main() {
	runner.Start()
}
