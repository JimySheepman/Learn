package main

// Prototype interface
//
// The Prototype interface declares the cloning methods.
// In most cases, it’s a single clone method.
type Inode interface {
	print(string)
	clone() Inode
}
