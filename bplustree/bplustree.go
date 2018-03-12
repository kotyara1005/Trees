package btree

import (
	"fmt"
	"sort"

	"github.com/kotyara1005/trees/bplustree/node"
)

// TODO add not uniq index

// BPlusTree class
type BPlusTree struct {
	root *node.Node
	t    int
	uniq bool
}

// Create new BTree
func Create(t int, uniq bool) *BPlusTree {
	x := node.Allocate(t)
	x.IsLeaf = true
	x.N = 0
	node.Write(x)
	tree := new(BPlusTree)
	tree.root = x
	tree.t = t
	tree.uniq = uniq
	return tree
}

// GetMaxKeysCount return double t
func (tree *BPlusTree) GetMaxKeysCount() int {
	return tree.t * 2
}

// GetSplitOffset return t + 1
func (tree *BPlusTree) GetSplitOffset() int {
	return tree.t + 1
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

// Returns delimiter, new Node
func (tree *BPlusTree) splitNode(n *node.Node, key int, link *node.Node) (int, *node.Node) {
	pos := sort.Search(n.N, func(i int) bool {
		return key <= n.Keys[i]
	})
	splitOffset := tree.GetSplitOffset()

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
	n.Next = newNode

	if pos < splitOffset {
		n.Insert(key, link, tree.uniq)
	} else {
		newNode.Insert(key, link, tree.uniq)
	}
	return newNode.Keys[0], newNode
}

func (tree *BPlusTree) insert(n *node.Node, key int) (int, *node.Node) {
	// check node writes
	var newChild *node.Node
	if !n.IsLeaf {
		// TODO use bin search
		i := n.N - 1
		for i >= 0 && key < n.Keys[i] {
			i = i - 1
		}
		i = i + 1
		next := node.Read(n.Links[i])
		key, newChild = tree.insert(next, key)
	}

	if newChild != nil || n.IsLeaf {
		if n.N < tree.GetMaxKeysCount() {
			n.Insert(key, newChild, n.IsLeaf && tree.uniq)
		} else {
			return tree.splitNode(n, key, newChild)
		}
	}
	return key, nil
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
		i--
		if i < n.N && i >= 0 && key == n.Keys[i] {
			return n, i
		}
		return nil, -1
	}
	return search(node.Read(n.Links[i]), key)
}

func (tree *BPlusTree) getLeftest() *node.Node {
	current := tree.root
	for !current.IsLeaf {
		current = current.Links[0]
	}
	return current
}

// SearchRange find range of keys and put it to channel
func (tree *BPlusTree) SearchRange(ch chan int, left *int, right *int) {
	var currentNode *node.Node
	var i int
	if left == nil {
		currentNode = tree.getLeftest()
		i = 0
	} else {
		currentNode, i = search(tree.root, *left)
	}
	if currentNode == nil {
		return
	}
	for currentNode != nil {
		for i < currentNode.N {
			if right != nil && currentNode.Keys[i] >= *right {
				break
			}
			if left == nil || currentNode.Keys[i] >= *left {
				ch <- currentNode.Keys[i]
			}
			i++
		}
		i = 0
		currentNode = currentNode.Next
	}
}

// Print BTree
func (tree *BPlusTree) Print() {
	fmt.Println(*tree)
	PrintTree(tree.root)
}

// PrintTree print B+ tree from node
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
