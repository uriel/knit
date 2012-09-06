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
	stitches["k"] = KnitStitch
	stitches["p"] = PurlStitch
}

// getStitchKind returns the kind of stitch represented by the
// supplied string.
func getStitchKind(s string) StitchKind {
	for k, v := range stitches {
		if len(s) > len(k) {
			continue
		}

		if !strings.HasPrefix(k, s) {
			continue
		}

		return v
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	line int
	col  int
	Kind StitchKind
}

func (s *Stitch) Line() int { return s.line }
func (s *Stitch) Col() int  { return s.col }
