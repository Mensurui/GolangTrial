package main

import (
	"io"
	"os"
)

var w io.Writer

func main() {
	w = os.Stdout
}
