package main

import "fmt"

// Concrete Prototype
//
// The Concrete Prototype class implements the cloning method. In addition to
// copying the original object’s data to the clone, this method may also handle
// some edge cases of the cloning process related to cloning linked objects,
// untangling recursive dependencies, etc.
type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{
		name: f.name + "_clone",
	}

	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren

	return cloneFolder
}
