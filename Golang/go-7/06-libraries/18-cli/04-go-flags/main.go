package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

		Offset uint `long:"offset" description:"Offset"`

		Call func(string) `short:"c" description:"Call phone number"`

		Name string `short:"n" long:"name" description:"A name" required:"true"`

		Animal string `long:"animal" choice:"cat" choice:"dog"`

		File string `short:"f" long:"file" description:"A file" value-name:"FILE"`

		Ptr *int `short:"p" description:"A pointer to an integer"`

		StringSlice []string `short:"s" description:"A slice of strings"`

		PtrSlice []*string `long:"ptrslice" description:"A slice of pointers to string"`

		IntMap map[string]int `long:"intmap" description:"A map from string to int"`
	}

	opts.Call = func(num string) {
		cmd := exec.Command("open", "callto:"+num)
		cmd.Start()
		cmd.Process.Release()
	}

	args := []string{
		"-vv",
		"--offset=5",
		"-n", "Me",
		"--animal", "dog",
		"-p", "3",
		"-s", "hello",
		"-s", "world",
		"--ptrslice", "hello",
		"--ptrslice", "world",
		"--intmap", "a:1",
		"--intmap", "b:5",
		"arg1",
		"arg2",
		"arg3",
	}

	args, err := flags.ParseArgs(&opts, args)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("Offset: %d\n", opts.Offset)
	fmt.Printf("Name: %s\n", opts.Name)
	fmt.Printf("Animal: %s\n", opts.Animal)
	fmt.Printf("Ptr: %d\n", *opts.Ptr)
	fmt.Printf("StringSlice: %v\n", opts.StringSlice)
	fmt.Printf("PtrSlice: [%v %v]\n", *opts.PtrSlice[0], *opts.PtrSlice[1])
	fmt.Printf("IntMap: [a:%v b:%v]\n", opts.IntMap["a"], opts.IntMap["b"])
	fmt.Printf("Remaining args: %s\n", strings.Join(args, " "))
}
