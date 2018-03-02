package utils

import (
	"fmt"

	"github.com/kotyara1005/trees/btree/node"
)

type Tree interface {
	GetT() int
	Insert(key int)
	Search(key int)
}

type Node interface {
	AddKey()
}

func PrintTree(n *node.Node) {
	fmt.Println(n)
	for i := 0; i < n.N; i++ {
		fmt.Println(n.Keys[i])
	}
	if !n.IsLeaf {
		for i := 0; i <= n.N; i++ {
			PrintTree(n.Links[i])
		}
	}
}
