// This program is to practice the basics of a Binary Tree data structure
package main

import (
	"binarytrees/binarytree"
	"fmt"
)

func main() {

	// a := binarytree.Build123a()
	// fmt.Print("a: ")
	// a.PrintTree()
	// fmt.Println()

	// b := binarytree.Build123b()
	// fmt.Print("b: ")
	// b.PrintTree()
	// fmt.Println()

	// c := binarytree.Build123c()
	// fmt.Print("c: ")
	// c.PrintTree()
	// fmt.Println()

	var a *binarytree.Node
	fmt.Println(a.IsEmpty())
	a = binarytree.Insert(a, 5)
	a = binarytree.Insert(a, 3)
	a = binarytree.Insert(a, 8)
	a = binarytree.Insert(a, 2)
	a = binarytree.Insert(a, 1)
	a = binarytree.Insert(a, 7)
	a = binarytree.Insert(a, 12)
	a = binarytree.Insert(a, 10)
	a = binarytree.Insert(a, 9)
	a.PrintTree()
	fmt.Println()
	fmt.Println("size:", a.Size())
	fmt.Println("Max Depth:", a.MaxDepth())
}
