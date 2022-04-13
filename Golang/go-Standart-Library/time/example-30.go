package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.UnixDate, "Sat Mar 7 11:06:39 PST 2015")
	if err != nil {
		panic(err)
	}
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}
	do("Unix", time.UnixDate, "Sat Mar  7 11:06:39 PST 2015")
	do("No pad", "<2>", "<7>")
	do("Spaces", "<_2>", "< 7>")
	do("Zeros", "<02>", "<07>")
	do("Suppressed pad", "04:05", "06:39")
}
