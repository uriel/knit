// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Group is a collection of nodes.
type Group struct {
	parent NodeCollection
	nodes  []Node
	line   int
	col    int
}

func (g *Group) Line() int              { return g.line }
func (g *Group) Col() int               { return g.col }
func (g *Group) Len() int               { return len(g.nodes) }
func (g *Group) Parent() NodeCollection { return g.parent }
func (g *Group) Append(n Node)          { g.nodes = append(g.nodes, n) }
func (g *Group) Nodes() []Node          { return g.nodes }

func (g *Group) Node(i int) Node {
	if i < 0 || i >= len(g.nodes) {
		return nil
	}
	return g.nodes[i]
}
