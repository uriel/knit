// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// A row node determines that the following pattern nodes
// belong to a given row. The row can optionally have a number
// defined for it.
type Row struct {
	Value int
	line  int
	col   int
}

// Line returns the original pattern source line number for this node.
func (r *Row) Line() int { return r.line }

// Col returns the original pattern source column number for this node.
func (r *Row) Col() int { return r.col }
