// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"strconv"
)

// Pattern represents a single, complete knitting pattern.
type Pattern struct {
	*Group        // Root node for the pattern's node tree.
	Name   string // Name of the pattern.
}

// MustParse parses the input pattern.
// It panics if an error occurred.
func MustParse(name, pat string) *Pattern {
	p, err := Parse(name, pat)

	if err != nil {
		panic(err)
	}

	return p
}

// Parse parses the given input pattern.
func Parse(name, pat string) (*Pattern, error) {
	p := new(Pattern)
	p.Name = name
	p.Group = new(Group)
	node := NodeCollection(p.Group)
	tokens := lex(pat)

loop:
	for {
		select {
		case tok := <-tokens:
			if tok == nil || tok.Type == tokEof {
				break loop
			}

			switch tok.Type {
			case tokError:
				return nil, fmt.Errorf("%s:%d:%d %s",
					name, tok.Line, tok.Col, tok.Data)

			case tokGroupStart:
				g := new(Group)
				g.line = tok.Line
				g.col = tok.Col
				g.parent = node
				node.Append(g)
				node = g

			case tokGroupEnd:
				node = node.Parent()

			case tokStitch:
				node.Append(&Stitch{
					tok.Line, tok.Col, getStitchKind(tok.Data),
				})

			case tokReference:
				node.Append(&Reference{tok.Data[1:], tok.Line, tok.Col})

			case tokNumber:
				n, err := strconv.ParseInt(tok.Data, 10, 32)

				if err != nil {
					return nil, fmt.Errorf("%s:%d:%d Invalid number %q,",
						name, tok.Line, tok.Col, tok.Data)
				}

				// A number can not directly follow another number.
				if sz := node.Len(); sz > 0 {
					switch node.Node(sz - 1).(type) {
					case *Number:
						return nil, fmt.Errorf(
							"%s:%d:%d Expected Stitch or Group, found Number %q,",
							name, tok.Line, tok.Col, tok.Data)
					}
				}

				node.Append(&Number{int(n), tok.Line, tok.Col})
			}
		}
	}

	return p, nil
}

// Expand uses the supplied handler to replace any external references
// with their actual data. It expands the referenced patterns recursively.
func (p *Pattern) Expand(rh ReferenceHandler) error {
	if rh == nil {
		return fmt.Errorf("Expand %q: Invalid reference handler.", p.Name)
	}

	err := expand(p, rh)

	if err != nil {
		return fmt.Errorf("Expand %q: %v", p.Name, err)
	}

	return nil
}

// Unroll unrolls all 'loop' constructs.
func (p *Pattern) Unroll() { unroll(p) }

// expand recursively expands pattern references.
func expand(list NodeCollection, rh ReferenceHandler) error {
	for i, node := range list.Nodes() {
		switch tt := node.(type) {
		case NodeCollection:
			err := expand(tt, rh)

			if err != nil {
				return err
			}

		case *Reference:
			ref, err := rh(tt.Name)
			if err != nil {
				return err
			}

			err = ref.Expand(rh)
			if err != nil {
				return err
			}

			list.SetNode(i, ref.Group)
		}
	}

	return nil
}

// unroll recursively unwinds loops.
func unroll(list NodeCollection) {
	var tmp []Node
	var elem Node
	var i, k int

	nodes := list.Nodes()

	for i = 0; i < len(nodes); i++ {
		switch tt := nodes[i].(type) {
		case NodeCollection:
			unroll(tt)

		case *Number:
			if i == 0 {
				continue
			}

			elem = nodes[i-1]
			tmp = make([]Node, tt.Value-1+len(nodes)-1)

			copy(tmp, nodes[:i])
			copy(tmp[i+tt.Value-1:], nodes[i+1:])

			// Repeat the previous element num - 1 times.
			for k = 0; k < tt.Value-1; k++ {
				tmp[i+k] = elem
			}

			nodes = tmp
		}
	}

	// Unpack groups.
	for i = 0; i < len(nodes); i++ {
		tt, ok := nodes[i].(NodeCollection)

		if !ok {
			continue
		}

		tmp = make([]Node, tt.Len()+len(nodes)-1)

		copy(tmp, nodes[:i])
		copy(tmp[i:], tt.Nodes())
		copy(tmp[i+tt.Len():], nodes[i+1:])

		nodes = tmp
	}

	list.SetNodes(nodes)
}
