package node

type Node struct {
	IsLeaf bool
	N      int
	// TODO customize
	Keys  []int
	Links []*Node
	Next  *Node
}

func Read(link *Node) *Node {
	// TODO read from disk
	return link
}

func Write(link *Node) {
	// TODO write to disk
}

func Allocate(t int) *Node {
	// t - minimum degree
	node := new(Node)
	node.Keys = make([]int, 2*t)
	node.Links = make([]*Node, 2*t+1)
	return node
}
