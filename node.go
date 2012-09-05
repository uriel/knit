// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

type Node interface {
	Line() int
	Col() int
}

type NodeCollection interface {
	Node
	Parent() NodeCollection
	Len() int
	Nodes() []Node
	Node(int) Node
	Append(Node)
}
