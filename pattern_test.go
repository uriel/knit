// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

var stdout = os.Stdout

func TestPattern(t *testing.T) {
	tests := []string{
		`Row 1 co 9 [ p 3 s 3 k 3 inc foo ] 10 bo 9`,
		`Row 2 co9[p3s3k3incfoo]10bo9`,
	}

	for i, str := range tests {
		_, err := Parse(fmt.Sprintf("Pattern %d", i), str)

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestExpand(t *testing.T) {
	a, err := Parse("patA", "[P3 K3] 10")
	if err != nil {
		t.Fatal(err)
	}

	b, err := Parse("patB", "P10 patA 10 P10")
	if err != nil {
		t.Fatal(err)
	}

	err = b.Expand(func(name string) (*Pattern, error) {
		if !strings.EqualFold(name, "patA") {
			t.Fatalf(`Expected reference "patA", got %q`, name)
		}

		return a, nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnroll(t *testing.T) {
	stitches := []StitchKind{
		PurlStitch,
		PurlStitch,
		KnitStitch,
		KnitStitch,
		Increase,
		KnitStitch,
		KnitStitch,
		Increase,
		KnitStitch,
		KnitStitch,
		Increase,
		KnitStitch,
		KnitStitch,
		Increase,
		PurlStitch,
		PurlStitch,
	}

	// These should all yield the exact same sequence of nodes after unrolling.
	a, err := Parse("PatternA", "P2 [K2 INC] 4 P2")
	if err != nil {
		t.Fatal(err)
	}

	b, err := Parse("PatternB", "P2 [[K2 INC] 2] 2 P2")
	if err != nil {
		t.Fatal(err)
	}

	c, err := Parse("PatternC", "P2 K2 INC K2 INC K2 INC K2 INC P2")
	if err != nil {
		t.Fatal(err)
	}

	d, err := Parse("PatternD", "P P K K INC K K INC K K INC K K INC P P")
	if err != nil {
		t.Fatal(err)
	}

	compareUnroll(t, a, stitches)
	compareUnroll(t, b, stitches)
	compareUnroll(t, c, stitches)
	compareUnroll(t, d, stitches)
}

func compareUnroll(t *testing.T, p *Pattern, stitches []StitchKind) {
	p.Unroll()

	if p.Len() != len(stitches) {
		t.Fatalf("%s len: Want %d, have %d", p.Name, len(stitches), p.Len())
	}

	nodes := p.Nodes()

	for i, node := range nodes {
		st, ok := node.(*Stitch)

		if !ok {
			t.Fatalf("%s:%d:%d type mismatch: Expected Stitch, have %T",
				p.Name, node.Line(), node.Col(), node)
		}

		if st.Kind != stitches[i] {
			t.Fatalf("%s:%d:%d Stitch mismatch: Expected %d, have %d",
				p.Name, node.Line(), node.Col(), stitches[i], st.Kind)
		}
	}
}

// dump writes a human-readable form of the pattern node tree
// to the given writer.
func dump(p *Pattern, w io.Writer) {
	if p.Group == nil || p.Group.Len() == 0 {
		fmt.Fprintf(w, "Pattern %q: <empty>\n", p.Name)
	} else {
		fmt.Fprintf(w, "Pattern %q:\n", p.Name)
	}

	dumpNodes(w, p.Group.Nodes(), " ")
}

// dumpNodes recursively writes nodes out to the given writer in
// a human-readable form.
func dumpNodes(w io.Writer, list []Node, indent string) {
	for _, node := range list {
		switch tt := node.(type) {
		case NodeCollection:
			fmt.Fprintf(w, "%s%03d:%03d %T {\n",
				indent, tt.Line(), tt.Col(), tt)
			dumpNodes(w, tt.Nodes(), indent+"  ")
			fmt.Fprintf(w, "%s}\n", indent)

		case *Stitch:
			var s string

			switch tt.Kind {
			case KnitStitch:
				s = "Knit"
			case PurlStitch:
				s = "Purl"
			case SlipStitch:
				s = "Slip"
			case CastOn:
				s = "CastOn"
			case BindOff:
				s = "BindOff"
			case Increase:
				s = "Increase"
			case Decrease:
				s = "Decrease"
			case YarnOver:
				s = "YarnOver"
			}

			fmt.Fprintf(w, "%s%03d:%03d %T(%q)\n",
				indent, tt.Line(), tt.Col(), tt, s)

		case *Row:
			fmt.Fprintf(w, "%s%03d:%03d %T(%d)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Value)

		case *Reference:
			fmt.Fprintf(w, "%s%03d:%03d %T(%q)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Name)

		case *Number:
			fmt.Fprintf(w, "%s%03d:%03d %T(%d)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Value)
		}
	}
}
