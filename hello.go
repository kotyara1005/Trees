package main

import (
	"fmt"

	btree "github.com/kotyara1005/trees/bplustree"
	// "github.com/kotyara1005/trees/btree"
)

func main() {
	t := btree.Create(2, true)
	for i := 1; i <= 15; i++ {
		t.Insert(i)
	}
	fmt.Println("==============")
	t.Print()
	fmt.Println("==============")

	ch := make(chan int)
	go func() {
		r := 9
		t.SearchRange(ch, nil, &r)
		close(ch)
	}()
	for {
		val, ok := <-ch
		if !ok {
			return
		}
		fmt.Println(val)
	}
}

func testSearch() {
	t := btree.Create(2, false)
	for i := 1; i <= 13; i++ {
		t.Insert(i)
	}
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
}
