// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Group is a group of nodes.
type Group struct {
	parent NodeCollection
	nodes  []Node
}

func (g *Group) Len() int               { return len(g.nodes) }
func (g *Group) Parent() NodeCollection { return g.parent }
func (g *Group) Append(v Node)          { g.nodes = append(g.nodes, v) }
func (g *Group) Nodes() []Node          { return g.nodes }

func (g *Group) Node(i int) Node {
	if i < 0 || i >= len(g.nodes) {
		return nil
	}
	return g.nodes[i]
}
