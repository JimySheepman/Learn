package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned")
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()
}
