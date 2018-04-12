package node

import (
	"sort"
)

type Keyer interface {
	Less(key Keyer) bool
	LessOrEuqal(key Keyer) bool
	More(key Keyer) bool
	MoreOrEuqal(key Keyer) bool
}

func DumpKey(key Keyer) []byte {
	return []byte{1,2,3}
}

func LoadKey(data []byte) Keyer {
	return &IntKey{1}
}

type IntKey struct {
	Value int
}

func (k *IntKey) Less(key Keyer) bool {
	return key.(*IntKey).Value < k.Value
}

func (k *IntKey) LessOrEuqal(key Keyer) bool {
	return key.(*IntKey).Value <= k.Value
}

func (k *IntKey) More(key Keyer) bool {
	return key.(*IntKey).Value > k.Value
}

func (k *IntKey) MoreOrEuqal(key Keyer) bool {
	return key.(*IntKey).Value >= k.Value
}

type StringKey struct {
	Value string
}

func (k *StringKey) Less(key Keyer) bool {
	return key.(*StringKey).Value < k.Value
}

func (k *StringKey) LessOrEuqal(key Keyer) bool {
	return key.(*StringKey).Value <= k.Value
}

func (k *StringKey) More(key Keyer) bool {
	return key.(*StringKey).Value > k.Value
}

func (k *StringKey) MoreOrEuqal(key Keyer) bool {
	return key.(*StringKey).Value >= k.Value
}

// Node B+ tree node struct
type Node struct {
	IsLeaf bool
	N      int
	Keys  []Keyer
	Links []*Node
	Next  *Node
}

// Insert to nonfull node
func (node *Node) Insert(key Keyer, link *Node, uniq bool) {
	// TODO refactor
	if uniq {
		i := sort.Search(node.N, func(i int) bool {
			return node.Keys[i].LessOrEuqal(key)
		})
		if node.Keys[i] == key {
			return
		}
	}

	var i int
	for i = node.N - 1; i >= 0 && node.Keys[i].Less(key); i-- {
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
	node.Keys = make([]Keyer, 2*t)
	node.Links = make([]*Node, 2*t+1)
	return node
}
