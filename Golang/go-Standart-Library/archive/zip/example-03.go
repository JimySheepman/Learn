package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
)

func main() {

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
}
