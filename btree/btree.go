package btree

import (
	"fmt"
	"sort"

	"github.com/kotyara1005/trees/btree/node"
)

// TODO improve search
// TODO add not uniq index
// TODO implement B+
// TODO fix worse case

type BTree struct {
	root *node.Node
	t    int
}

func Create(t int) *BTree {
	x := node.Allocate(t)
	x.IsLeaf = true
	x.N = 0
	node.Write(x)
	tree := new(BTree)
	tree.root = x
	tree.t = t
	return tree
}

func (tree *BTree) Insert(key int) {
	if tree.root.N == 2*tree.t-1 {
		root := tree.root
		tree.root = node.Allocate(tree.t)
		tree.root.IsLeaf = false
		tree.root.N = 0
		tree.root.Links[0] = root
		tree.SplitChild(tree.root, 0)
	}
	tree.InsertNonfull(tree.root, key)
}

func (tree *BTree) SplitChild(parent *node.Node, i int) {
	newNode := node.Allocate(tree.t)
	child := parent.Links[i]
	newNode.IsLeaf = child.IsLeaf
	newNode.N = tree.t - 1
	for j := 0; j < tree.t-1; j++ {
		newNode.Keys[j] = child.Keys[j+tree.t]
	}
	if !child.IsLeaf {
		for j := 0; j < tree.t; j++ {
			newNode.Links[j] = child.Links[j+tree.t]
		}
	}
	child.N = tree.t - 1
	for j := parent.N + 1; j > i+1; j-- {
		parent.Links[j+1] = parent.Links[j]
	}
	parent.Links[i+1] = newNode
	for j := parent.N; j > i; j-- {
		parent.Keys[j+1] = parent.Keys[j]
	}
	parent.Keys[i] = child.Keys[tree.t-1]
	parent.N += 1
	node.Write(child)
	node.Write(newNode)
	node.Write(parent)
}

func (tree *BTree) InsertNonfull(n *node.Node, key int) {
	i := n.N - 1
	if n.IsLeaf {
		fmt.Println(n)
		for i >= 0 && key < n.Keys[i] {
			n.Keys[i+1] = n.Keys[i]
			i -= 1
		}
		// fmt.Println(i, *n)
		n.Keys[i+1] = key
		n.N += 1
		node.Write(n)
	} else {
		for i >= 0 && key < n.Keys[i] {
			i = i - 1
		}
		i = i + 1
		next := node.Read(n.Links[i])
		fmt.Println(next)
		fmt.Println(i)
		if next.N == 2*tree.t-1 {
			tree.SplitChild(n, i)
			if key > n.Keys[i] {
				i = i + 1
				next = node.Read(n.Links[i])
			}
		}
		fmt.Println(n, i)
		tree.InsertNonfull(next, key)
	}
}

func (tree *BTree) Search(key int) (*node.Node, int) {
	return search(tree.root, key)
}

func search(n *node.Node, key int) (*node.Node, int) {
	i := sort.Search(n.N, func(i int) bool {
		return n.Keys[i] >= key
	})
	if i < n.N && key == n.Keys[i] {
		return n, i
	} else if n.IsLeaf {
		return nil, -1
	} else {
		return search(node.Read(n.Links[i]), key)
	}
}

func (tree *BTree) Print() {
	fmt.Println(*tree)
	printTree(tree.root)
}

func printTree(n *node.Node) {
	fmt.Println(n)
	for i := 0; i < n.N; i++ {
		fmt.Println(n.Keys[i])
	}
	if !n.IsLeaf {
		for i := 0; i <= n.N; i++ {
			printTree(n.Links[i])
		}
	}
}

func Remove() {
	// TODO
}

func Rebuild() {
	// TODO
}
