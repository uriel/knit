// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPattern(t *testing.T) {
	tests := []string{
		`co 9 [ p 3 s 3 k 3 inc $foo ] 10 bo +`,
		`co9[p3s3k3inc$foo]10bo+`,
	}

	for i, str := range tests {
		p, err := Parse(fmt.Sprintf("Pattern %d", i), str)

		if err != nil {
			t.Fatal(err)
		}

		dump(p, os.Stdout)
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

	dumpNodes(w, p.Group, " ")
}

// dumpNodes recursively writes nodes out to the given writer in
// a human-readable form.
func dumpNodes(w io.Writer, list NodeCollection, indent string) {
	for _, node := range list.Nodes() {
		switch tt := node.(type) {
		case NodeCollection:
			fmt.Fprintf(w, "%s%03d:%03d %T {\n",
				indent, tt.Line(), tt.Col(), tt)
			dumpNodes(w, tt, indent+"  ")
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

		case *Quantifier:
			fmt.Fprintf(w, "%s%03d:%03d %T(%q)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Kind)

		case *Reference:
			fmt.Fprintf(w, "%s%03d:%03d %T(%q)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Name)

		case *Number:
			fmt.Fprintf(w, "%s%03d:%03d %T(%d)\n",
				indent, tt.Line(), tt.Col(), tt, tt.Value)
		}
	}
}
