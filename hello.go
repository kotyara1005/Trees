package main

import (
	"fmt"

	btree "github.com/kotyara1005/trees/bplustree"
	// "github.com/kotyara1005/trees/btree"
)

func main() {
	// fmt.Printf("hello, world\n")
	t := btree.Create(2)
	// t.Print()
	// fmt.Println("==============")
	for i := 1; i <= 5; i++ {
		t.Insert(i)
	}
	// t.Insert(2)
	fmt.Println("==============")
	t.Print()
	fmt.Println("==============")
	fmt.Println(t.Search(3))
	// fmt.Println("==============")
	// fmt.Println(t.Search(9))
	// fmt.Println("==============")
}
