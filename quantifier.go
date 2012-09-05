// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

type QuantifierKind uint8

// Known quantifier kinds.
const (
	UnknownQuantifier QuantifierKind = iota
	FillRowQuantifier
)

func (s QuantifierKind) String() string {
	switch s {
	case FillRowQuantifier:
		return "FillRow"
	}

	panic("unreachable")
}

func getQuantifierKind(v string) QuantifierKind {
	switch v {
	case "+":
		return FillRowQuantifier
	}

	return UnknownQuantifier
}

// A quantifier holds an abstract representation of the amount
// of times to repeat the preceeding group or stitch.
//
// For example: `P+`.
//
// The `+` symbol is a quantifier which tells us the P stitch should
// be repeated until the end of the row, regardless of how many times
// that might be.
type Quantifier struct {
	Kind QuantifierKind
}
