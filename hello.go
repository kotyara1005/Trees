package main

import (
	"fmt"

	btree "github.com/kotyara1005/trees/bplustree"
	// "github.com/kotyara1005/trees/btree"
)

func main() {
	// key := 3
	// arr := [5]int{1,2,4,5,6}
	// i := sort.Search(5, func(i int) bool {
	// 	return key <= arr[i]
	// })
	// fmt.Println(i)

	// fmt.Printf("hello, world\n")
	t := btree.Create(2)
	// t.Print()
	// fmt.Println("==============")
	for i := 1; i <= 13; i++ {
		t.Insert(i)
	}
	// t.Insert(2)
	fmt.Println("==============")
	t.Print()
	fmt.Println("==============")
	fmt.Println(t.Search(1))
	fmt.Println("==============")
	fmt.Println(t.Search(3))
	fmt.Println("==============")
	fmt.Println(t.Search(7))
	fmt.Println("==============")
	fmt.Println(t.Search(11))
	fmt.Println("==============")
	fmt.Println(t.Search(13))
	// fmt.Println("==============")
	// fmt.Println(t.Search(9))
	// fmt.Println("==============")
}
