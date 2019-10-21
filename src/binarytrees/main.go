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
	a = binarytree.Insert(a, 7)
	a = binarytree.Insert(a, 12)
	a = binarytree.Insert(a, 10)
	a = binarytree.Insert(a, 9)
	a = binarytree.Insert(a, 13)
	fmt.Print("Inorder traversal: ")
	a.PrintTree()
	fmt.Println()
	fmt.Print("Post order traversal: ")
	a.PrintPostorder()
	fmt.Println()
	fmt.Println("size:", a.Size())
	fmt.Println("Max Depth:", a.MaxDepth())
	fmt.Println("Min Value:", a.MinValue())
	sum := 10
	fmt.Printf("HasPathSum (%d): %t\n", sum, a.HasPathSum(sum))
	fmt.Println("Tree Paths: ")
	a.PrintPaths()
	mirror := a.Mirror()
	fmt.Print("Mirrored Tree: ")
	mirror.PrintTree()
	fmt.Println()
	// a.DoubleTree()
	// a.PrintTree()
	var c *binarytree.Node
	c = binarytree.Insert(c, 5)
	c = binarytree.Insert(c, 3)
	c = binarytree.Insert(c, 8)
	fmt.Println(binarytree.SameTree(a, c))
}
