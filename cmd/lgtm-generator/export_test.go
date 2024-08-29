package main

import (
	"io"
)

func Run(out io.Writer, tag string) error {
	return (&giphyRandomCmd{}).run(out, tag)
}
