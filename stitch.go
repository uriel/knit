// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

type StitchKind uint8

// Known stitch kinds.
const (
	KnitStitch StitchKind = iota
	PurlStitch
	SlipStitch
	CastOn
	BindOff
	Increase
	Decrease
	YarnOver
)

func (s StitchKind) String() string {
	switch s {
	case KnitStitch:
		return "K"
	case PurlStitch:
		return "P"
	case SlipStitch:
		return "S"
	case CastOn:
		return "CO"
	case BindOff:
		return "BO"
	case Increase:
		return "INC"
	case Decrease:
		return "DEC"
	case YarnOver:
		return "YO"
	}

	panic("unreachable")
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	line int
	col  int
	Kind StitchKind
}

// newStitch creates a new stitch from the given type.
// Returns nil if the type is not recognized as a valid stitch.
// We expect v to be something like 'k' for Knit stitch or 'p'
// for Purl stitch.
func newStitch(v string, line, col int) *Stitch {
	s := new(Stitch)
	s.line = line
	s.col = col

	switch v {
	case "k":
		s.Kind = KnitStitch
	case "p":
		s.Kind = PurlStitch
	case "sl", "s":
		s.Kind = SlipStitch
	case "co", "c":
		s.Kind = CastOn
	case "bo", "b":
		s.Kind = BindOff
	case "inc":
		s.Kind = Increase
	case "dec", "tog":
		s.Kind = Decrease
	case "yo":
		s.Kind = YarnOver

	default:
		return nil
	}

	return s
}

func (s *Stitch) Line() int { return s.line }
func (s *Stitch) Col() int  { return s.col }
