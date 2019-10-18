// This program is to practice the basics of a Binary Tree data structure
package main

import (
	"binarytrees/binarytree"
	"fmt"
)

func main() {

	// binarytree.Build123a()

	var a *binarytree.Node

	fmt.Println(a.IsEmpty())

	a = binarytree.Insert(a, 5)
	a = binarytree.Insert(a, 3)
	a = binarytree.Insert(a, 8)

	a.PrintTree()
}
