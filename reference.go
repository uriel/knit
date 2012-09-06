// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// A ReferenceHandler is called by the parser when
// it is instructed to expand all reference nodes.
// The handler should return the compiled pattern for
// the given name.
type ReferenceHandler func(name string) (*Pattern, error)

// A Reference holds the name of an external pattern we are
// referencing in the current pattern.
//
// References are not expanded by the pattern parser.
// This must be done at a later stage by the host application.
// We therefore do not validate the existence of a referenced pattern.
type Reference struct {
	Name string
	line int
	col  int
}

// Line returns the original pattern source line number for this node.
func (r *Reference) Line() int { return r.line }

// Col returns the original pattern source column number for this node.
func (r *Reference) Col() int { return r.col }
