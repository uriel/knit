// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import "strings"

type StitchKind uint8

// Known stitch kinds.
const (
	UnknownStitch StitchKind = iota
	KnitStitch
	PurlStitch
	SlipStitch
	CastOn
	KnitOn
	PurlOn
	BindOff
	Increase
	Decrease
	YarnOver
)

// Listing of known stitches.
var stitches map[string]StitchKind

func init() {
	stitches = make(map[string]StitchKind)
	stitches["inc"] = Increase
	stitches["dec"] = Decrease
	stitches["tog"] = Decrease
	stitches["yo"] = YarnOver
	stitches["sl"] = SlipStitch
	stitches["s"] = SlipStitch
	stitches["co"] = CastOn
	stitches["bo"] = BindOff
	stitches["ko"] = KnitOn
	stitches["po"] = PurlOn
	stitches["k"] = KnitStitch
	stitches["p"] = PurlStitch
}

// getStitchKind returns the kind of stitch represented by the
// supplied string.
func getStitchKind(s string) StitchKind {
	for k, v := range stitches {
		if strings.EqualFold(k, s) {
			return v
		}
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	line int
	col  int
	Kind StitchKind
}

// Line returns the original pattern source line number for this node.
func (s *Stitch) Line() int { return s.line }

// Col returns the original pattern source column number for this node.
func (s *Stitch) Col() int { return s.col }
