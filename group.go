// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Group is a collection of nodes.
type Group struct {
	parent *Group
	nodes  []Node
	line   int
	col    int
}

// Line returns the original pattern source line number for this node.
func (g *Group) Line() int { return g.line }

// Col returns the original pattern source column number for this node.
func (g *Group) Col() int { return g.col }

// Len returns the length of the node list.
func (g *Group) Len() int { return len(g.nodes) }

// Pattern returns the group's parent node.
func (g *Group) Parent() *Group { return g.parent }

// Nodes returns the list of nodes for this group.
func (g *Group) Nodes() []Node { return g.nodes }

// SetNodes sets the nodes to the gives list.
func (g *Group) SetNodes(list []Node) { g.nodes = list }

// Append appends the given set of nodes to the group.
func (g *Group) Append(argv ...Node) {
	g.nodes = append(g.nodes, argv...)
}

// Node returns the node at the given index.
func (g *Group) Node(i int) Node {
	if i < 0 || i >= len(g.nodes) {
		return nil
	}
	return g.nodes[i]
}

// SetNode sets the node at the given index.
func (g *Group) SetNode(i int, n Node) {
	if i < 0 || i >= len(g.nodes) {
		return
	}
	g.nodes[i] = n
}
