// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// Node represents a single node.
type Node interface {
	Line() int
	Col() int
}
