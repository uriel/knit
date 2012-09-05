// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

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

func getStitchKind(v string) StitchKind {
	switch v {
	case "k":
		return KnitStitch
	case "p":
		return PurlStitch
	case "sl", "s":
		return SlipStitch
	case "co", "c":
		return CastOn
	case "bo", "b":
		return BindOff
	case "inc":
		return Increase
	case "dec", "tog":
		return Decrease
	case "yo":
		return YarnOver
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	Kind StitchKind
}
