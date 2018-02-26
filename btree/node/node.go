package node

// t - minimum degree

type Node struct {
	IsLeaf bool
	N      int
	Keys   []int
	Links  []*Node
	Values []int
}

func Read(link *Node) *Node {
	// TODO read from disk
	return link
}

func Write(link *Node) {
	// TODO write to disk
}

func Allocate(t int) *Node {
	node := new(Node)
	node.Keys = make([]int, 2*t-1)
	node.Links = make([]*Node, 2*t)
	return node
}
