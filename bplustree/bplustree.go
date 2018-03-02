package btree

import (
	"fmt"
	"sort"

	"github.com/kotyara1005/trees/bplustree/node"
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

// GetMaxKeysCount return double t
func (tree *BPlusTree) GetMaxKeysCount() int {
	return tree.t * 2
}

// Insert new key to BPlusTree
func (tree *BPlusTree) Insert(key int) {
	key, newChild := tree.insert(tree.root, key)
	if newChild != nil {
		oldRoot := tree.root
		tree.root = node.Allocate(tree.t)
		tree.root.IsLeaf = false
		tree.root.N = 1
		tree.root.Keys[0] = key
		tree.root.Links[0] = oldRoot
		tree.root.Links[1] = newChild
	}
}

func (tree *BPlusTree) splitChild1(parent *node.Node, i int) {
	// create new node
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

// Returns delimiter, new Node
func (tree *BPlusTree) splitNode(n *node.Node, key int, link *node.Node) (int, *node.Node) {
	fmt.Println(n)
	splitOffset := tree.t + 1

	pos := sort.Search(n.N, func(i int) bool {
		return key <= n.Keys[i]
	})
	if pos < splitOffset {
		splitOffset--
	}

	newNode := node.Allocate(tree.t)
	newNode.IsLeaf = n.IsLeaf
	newNode.N = n.N - splitOffset

	for i := 0; i < newNode.N; i++ {
		newNode.Keys[i] = n.Keys[splitOffset+i]
	}
	if !n.IsLeaf {
		for i := 0; i < newNode.N+1; i++ {
			newNode.Links[i] = n.Links[splitOffset+i]
		}
	}
	n.N = splitOffset

	if pos < splitOffset {
		n.Insert(key, link)
	} else {
		newNode.Insert(key, link)
	}
	fmt.Println(n, newNode)
	return newNode.Keys[0], newNode
}

func (tree *BPlusTree) insert(n *node.Node, key int) (int, *node.Node) {
	// check node writes
	i := n.N - 1
	if n.IsLeaf {
		if n.N == tree.GetMaxKeysCount() {
			return tree.splitNode(n, key, nil)
		} else {
			n.Insert(key, nil)
			return key, nil
		}
	} else {
		for i >= 0 && key < n.Keys[i] {
			i = i - 1
		}
		i = i + 1
		next := node.Read(n.Links[i])

		key, newChild := tree.insert(next, key)
		fmt.Println(key, newChild)
		if newChild != nil {
			if n.N == tree.GetMaxKeysCount() {
				return tree.splitNode(n, key, newChild)
			} else {
				n.Insert(key, newChild)
				return key, nil
			}
		}
		return key, nil
	}
}

// Search first key in BTree
func (tree *BPlusTree) Search(key int) (*node.Node, int) {
	return search(tree.root, key)
}

func search(n *node.Node, key int) (*node.Node, int) {
	i := sort.Search(n.N, func(i int) bool {
		return key < n.Keys[i]
	})
	if n.IsLeaf {
		fmt.Println(n)
		fmt.Println(i, key, n.Keys[i - 1])
		i--
		if i < n.N && i >= 0 && key == n.Keys[i] {
			return n, i
		}
		return nil, -1
	}
	return search(node.Read(n.Links[i]), key)
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
	PrintTree(tree.root)
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

// Remove key from BTree
func Remove() {
	// TODO
}
