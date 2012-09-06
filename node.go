// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Node represents a single node.
type Node interface {
	Line() int
	Col() int
}

// NodeCollection represents a group of nodes.
// This is itself a node.
type NodeCollection interface {
	Node
	Parent() NodeCollection
	Len() int
	SetNodes([]Node)
	Nodes() []Node
	SetNode(int, Node)
	Node(int) Node
	Append(...Node)
}
