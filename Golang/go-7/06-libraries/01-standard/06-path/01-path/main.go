package main

import (
	"fmt"
	"path"
)

func main() {
	Sample("Dir", Dir)
	Sample("Base", Base)
	Sample("Ext", Ext)
	Sample("Split", Split)
	Sample("Join", Join)
	Sample("Clean", Clean)
	Sample("Match", Match)
	Sample("IsAbs", IsAbs)
}

func Dir() {
	// Returns dir of path
	fn := func(p string) {
		dir := path.Dir(p)
		fmt.Printf("%q -> %q\n", p, dir)
	}
	fn("")
	fn("/home/merlins/Projects/Learn")
}

func Base() {
	// Returns last element of path
	fn := func(p string) {
		base := path.Base(p)
		fmt.Printf("%q -> %q\n", p, base)
	}
	fn("/home/merlins/Projects/Learn")
}

func Ext() {
	// Returns ext of path
	fn := func(p string) {
		ext := path.Ext(p)
		fmt.Printf("%q -> %q\n", p, ext)
	}

	fn("/home/merlins/Projects/Learn/main.go")
}

func Split() {
	// Returns split of path
	fn := func(p string) {
		dir, file := path.Split(p)
		fmt.Printf("%q -> %q, %q\n", p, dir, file)
	}
	fn("/home/merlins/Projects/Learn/main.go")
}

func Join() {
	// Joins paths
	fn := func(items ...string) {
		p := path.Join(items...)
		fmt.Printf("%q -> %q\n", items, p)
	}
	fn("/", "home", "merlins", "Projects", "Learn", "main.go.go")
}

func Clean() {
	// Clean paths
	fn := func(p string) {
		clean := path.Clean(p)
		fmt.Printf("%q -> %q\n", p, clean)
	}
	fn("/tmp/GoLand")
}

func Match() {
	// Matches path with pattern
	fn := func(pattern, name string) {
		matched, err := path.Match(pattern, name)
		if err != nil {
			fmt.Printf("%q, %q -> %v\n", name, pattern, err)
		} else {
			fmt.Printf("%q, %q -> %t\n", name, pattern, matched)
		}
	}
	fn("abc", "abc")
	fn("a*", "abc")
	fn("a*/b", "a/c/b")
}

func IsAbs() {
	// Checks path is absolute
	fn := func(p string) {
		isAbs := path.IsAbs(p)
		fmt.Printf("%q -> %t\n", p, isAbs)
	}
	fn("/tmp/GoLand")
	fn("c:/")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
