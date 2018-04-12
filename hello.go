package main

import (
	"fmt"

	btree "github.com/kotyara1005/trees/bplustree"
	node "github.com/kotyara1005/trees/bplustree/node"
	// "github.com/kotyara1005/trees/btree"
)

func main() {
	t := btree.Create(2, true)
	for i := 1; i <= 15; i++ {
		t.Insert(&node.IntKey{Value: i})
	}
	fmt.Println("==============")
	t.Print()
	fmt.Println("==============")

	ch := make(chan node.Keyer)
	go func() {
		r := node.IntKey{9}
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
		t.Insert(&node.IntKey{Value: i})
	}
	fmt.Println("==============")
	t.Print()
	fmt.Println("==============")
	fmt.Println(t.Search(&node.IntKey{Value: 1}))
	fmt.Println("==============")
	fmt.Println(t.Search(&node.IntKey{Value: 3}))
	fmt.Println("==============")
	fmt.Println(t.Search(&node.IntKey{Value: 7}))
	fmt.Println("==============")
	fmt.Println(t.Search(&node.IntKey{Value: 11}))
	fmt.Println("==============")
	fmt.Println(t.Search(&node.IntKey{Value: 13}))
}
