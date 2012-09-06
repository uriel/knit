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

				// A number can not follow a number or quantifier.
				if sz := node.Len(); sz > 0 {
					switch node.Node(sz - 1).(type) {
					case *Number, *Quantifier:
						return nil, fmt.Errorf(
							"%s:%d:%d Expected Stitch or Group, found Number %q,",
							name, tok.Line, tok.Col, tok.Data)
					}
				}

				node.Append(&Number{n, tok.Line, tok.Col})

			case tokQuantifier:
				q := newQuantifier(tok.Data, tok.Line, tok.Col)

				if q == nil {
					return nil, fmt.Errorf("%s:%d:%d Unknown quantifier kind %q,",
						name, tok.Line, tok.Col, tok.Data)
				}

				// A quantifier can not follow a number or quantifier.
				if sz := node.Len(); sz > 0 {
					switch node.Node(sz - 1).(type) {
					case *Number, *Quantifier:
						return nil, fmt.Errorf(
							"%s:%d:%d Expected Stitch or Group, found Quantifier %q,",
							name, tok.Line, tok.Col, tok.Data)
					}
				}

				node.Append(q)
			}
		}
	}

	return p, nil
}

// Expand uses the supplied handler to return a copy of this pattern
// with any references replaced by their actual data.
func (p *Pattern) Expand(rh ReferenceHandler) (*Pattern, error) {
	np := new(Pattern)
	np.Name = p.Name

	return np, nil
}

// Unwind unrolls all 'loop' constructs.
// It returns a copy of the original pattern.
func (p *Pattern) Unwind() (*Pattern, error) {
	np := new(Pattern)
	np.Name = p.Name

	return np, nil
}
