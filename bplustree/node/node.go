package node

import (
	"sort"
)

type Node struct {
	IsLeaf bool
	N      int
	// TODO customize
	Keys  []int
	Links []*Node
	Next  *Node
}

// Insert to nonfull node
func (node *Node) Insert(key int, link *Node, uniq bool) {
	// TODO refactor
	if uniq {
		i := sort.Search(node.N, func(i int) bool {
			return key <= node.Keys[i]
		})
		if node.Keys[i] == key {
			return
		}
	}

	var i int
	for i = node.N - 1; i >= 0 && key < node.Keys[i]; i-- {
		node.Keys[i+1] = node.Keys[i]
	}
	node.Keys[i+1] = key
	node.N++

	if !node.IsLeaf {
		var j int
		for j = node.N - 1; j > i+1; j-- {
			node.Links[j+1] = node.Links[j]
		}
		node.Links[j+1] = link
	}
}

// Read node from disk
func Read(link *Node) *Node {
	// TODO read from disk
	return link
}

// Write node to disk
func Write(link *Node) {
	// TODO write to disk
}

// Allocate new node
func Allocate(t int) *Node {
	// t - minimum degree
	node := new(Node)
	node.Keys = make([]int, 2*t)
	node.Links = make([]*Node, 2*t+1)
	return node
}
