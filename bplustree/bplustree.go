package btree

import (
	"fmt"
	"sort"

	"github.com/kotyara1005/trees/bplustree/node"
	"github.com/kotyara1005/trees/utils"
)

// TODO implement B+
// TODO add not uniq index
// TODO improve search
// TODO fix worse case

// BPlusTree class
type BPlusTree struct {
	root *node.Node
	t    int
}

// Create new BTree
func Create(t int) *BPlusTree {
	x := node.Allocate(t)
	x.IsLeaf = true
	x.N = 0
	node.Write(x)
	tree := new(BPlusTree)
	tree.root = x
	tree.t = t
	return tree
}

// Insert new key to BTree
func (tree *BPlusTree) Insert(key int) {
	if tree.root.N == 2*tree.t-1 {
		root := tree.root
		tree.root = node.Allocate(tree.t)
		tree.root.IsLeaf = false
		tree.root.N = 0
		tree.root.Links[0] = root
		tree.splitChild(tree.root, 0)
	}
	tree.insertNonfull(tree.root, key)
}

func (tree *BPlusTree) splitChild(parent *node.Node, i int) {
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
	child.N = tree.t
	for j := parent.N + 1; j > i+1; j-- {
		parent.Links[j+1] = parent.Links[j]
	}
	parent.Links[i+1] = newNode
	for j := parent.N; j > i; j-- {
		parent.Keys[j+1] = parent.Keys[j]
	}
	parent.Keys[i] = child.Keys[tree.t-1]
	parent.N++
	node.Write(child)
	node.Write(newNode)
	node.Write(parent)
}

func (tree *BPlusTree) insertNonfull(n *node.Node, key int) {
	i := n.N - 1
	if n.IsLeaf {
		for i >= 0 && key < n.Keys[i] {
			n.Keys[i+1] = n.Keys[i]
			i--
		}
		n.Keys[i+1] = key
		n.N++
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
			tree.splitChild(n, i)
			if key > n.Keys[i] {
				i = i + 1
				next = node.Read(n.Links[i])
			}
		}
		tree.insertNonfull(next, key)
	}
}

// Search first key in BTree
func (tree *BPlusTree) Search(key int) (*node.Node, int) {
	return search(tree.root, key)
}

func search(n *node.Node, key int) (*node.Node, int) {
	i := sort.Search(n.N, func(i int) bool {
		return key <= n.Keys[i]
	})
	if n.IsLeaf {
		if i < n.N && key == n.Keys[i] {
			return n, i
		} else {
			return nil, -1
		}
	} else {
		return search(node.Read(n.Links[i]), key)
	}
}

// func searchAll(n *node.Node, key int) (*node.Node, int) {
// 	i := sort.Search(n.N, func(i int) bool {
// 		return key <= n.Keys[i]
// 	})
// 	fmt.Println(i)
// 	if i < n.N && key == n.Keys[i] {
// 		fmt.Println(n, i)
// 		fmt.Println(node.Read(n.Links[i]))
// 		return search(node.Read(n.Links[i+1]), key, all)
// 	}
// 	if n.IsLeaf {
// 		return nil, -1
// 	} else {
// 		return search(node.Read(n.Links[i]), key, all)
// 	}
// }

// Print BTree
func (tree *BPlusTree) Print() {
	fmt.Println(*tree)
	utils.PrintTree(tree.root)
}

// Remove key from BTree
func Remove() {
	// TODO
}
