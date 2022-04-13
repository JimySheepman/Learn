package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil {
		panic(err)
	}

	fmt.Println("default format:", t)
	fmt.Println("Unix format:", t.Format(time.UnixDate))
	fmt.Println("Same, in UTC:", t.UTC().Format(time.UnixDate))
	do := func(name, layout, want string) {
		got := t.Format(layout)
		if want != got {
			fmt.Printf("error: for %q got %q; expected %q\n", layout, got, want)
			return
		}
		fmt.Printf("%-16s %q gives %q\n", name, layout, got)
	}

	fmt.Printf("\nFormats:\n\n")
	do("Basic full date", "Mon Jan 2 15:04:05 MST 2006", "Wed Feb 25 11:06:39 PST 2015")
	do("Basic short date", "2006/01/02", "2015/02/25")
	do("AM/PM", "3PM==3pm==15h", "11AM==11am==11h")

	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39.1234 PST 2015")
	if err != nil {
		panic(err)
	}

	do("No fraction", time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	do("0s for fraction", "15:04:05.00000", "11:06:39.12340")
	do("9s for fraction", "15:04:05.99999999", "11:06:39.1234")
}