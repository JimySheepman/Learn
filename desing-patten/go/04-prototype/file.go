package main

import "fmt"

// Concrete Prototype
//
// The Concrete Prototype class implements the cloning method. In addition to
// copying the original object’s data to the clone, this method may also handle
// some edge cases of the cloning process related to cloning linked objects,
// untangling recursive dependencies, etc.
type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}
