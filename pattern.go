// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"io"
	"strconv"
)

// Pattern represents a single, complete knitting pattern.
type Pattern struct {
	Name string         // Name of the pattern.
	Root NodeCollection // Root node for the pattern's node tree.
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
	p.Root = new(Group)
	node := p.Root
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
				g.parent = node
				node.Append(g)
				node = g

			case tokGroupEnd:
				node = node.Parent()

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

				node.Append(&Number{n})

			case tokQuantifier:
				kind := getQuantifierKind(tok.Data)

				if kind == UnknownQuantifier {
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

				node.Append(&Quantifier{kind})

			case tokStitch:
				kind := getStitchKind(tok.Data)

				if kind == UnknownStitch {
					return nil, fmt.Errorf("%s:%d:%d Unknown stitch kind %q,",
						name, tok.Line, tok.Col, tok.Data)
				}

				node.Append(&Stitch{kind})
			}
		}
	}

	return p, nil
}

// dump writes a human-readable form of the pattern node tree
// to the given writer. This is for debugging only.
func (p *Pattern) dump(w io.Writer) {
	if p.Root == nil || p.Root.Len() == 0 {
		fmt.Fprintf(w, "Pattern %q: <empty>\n", p.Name)
	} else {
		fmt.Fprintf(w, "Pattern %q:\n", p.Name)
	}

	dumpNodes(w, p.Root, " ")
}

// dumpNodes recursively dumps nodes out to the guven writer in
// a human-readable form. For debugging purposes only.
func dumpNodes(w io.Writer, list NodeCollection, indent string) {
	for _, node := range list.Nodes() {
		switch tt := node.(type) {
		case NodeCollection:
			fmt.Fprintf(w, "%s%T {\n", indent, tt)
			dumpNodes(w, tt, indent+"  ")
			fmt.Fprintf(w, "%s}\n", indent)

		case *Stitch:
			fmt.Fprintf(w, "%s%T(%q)\n", indent, tt, tt.Kind)

		case *Quantifier:
			fmt.Fprintf(w, "%s%T(%q)\n", indent, tt, tt.Kind)

		case *Number:
			fmt.Fprintf(w, "%s%T(%d)\n", indent, tt, tt.Value)
		}
	}
}
