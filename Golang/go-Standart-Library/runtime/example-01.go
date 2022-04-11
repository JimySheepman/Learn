package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	c := func() {
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			return
		}

		pc = pc[:n]
		frames := runtime.CallersFrames(pc)

		for {
			frame, more := frames.Next()

			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)

			if !more {
				break
			}
		}
	}

	b := func() { c() }
	a := func() { b() }

	a()
}
