package main

// Prototype interface
type Inode interface {
	print(string)
	clone() Inode
}
