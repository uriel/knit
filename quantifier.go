// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

type QuantifierKind uint8

// Known quantifier kinds.
const (
	FillRow QuantifierKind = iota
)

func (q QuantifierKind) String() string {
	switch q {
	case FillRow:
		return "FillRow"
	}

	panic("unreachable")
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
	line int
	col  int
	Kind QuantifierKind
}

// newQuantifier creates a new quantifier.
// Returns nil if v is not a valid quantifier kind.
func newQuantifier(v string, line, col int) *Quantifier {
	q := new(Quantifier)
	q.line = line
	q.col = col

	switch v {
	case "+":
		q.Kind = FillRow

	default:
		return nil
	}

	return q
}

func (q *Quantifier) Line() int { return q.line }
func (q *Quantifier) Col() int  { return q.col }
