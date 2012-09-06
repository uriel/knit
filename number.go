// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// A number holds a concrete amount of times to repeat the preceeding
// group or stitch.
//
// For example in the pattern `P3`, the number `3` tells us the P stitch
// should be repeated exactly three times.
type Number struct {
	Value int
	line  int
	col   int
}

// Line returns the original pattern source line number for this node.
func (n *Number) Line() int { return n.line }

// Col returns the original pattern source column number for this node.
func (n *Number) Col() int { return n.col }
