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
	CastOnStitch
	BindOffStitch
	IncreaseStitch
	DecreaseStitch
)

func (s StitchKind) String() string {
	switch s {
	case KnitStitch:
		return "K"
	case PurlStitch:
		return "P"
	case SlipStitch:
		return "S"
	case CastOnStitch:
		return "CO"
	case BindOffStitch:
		return "BO"
	case IncreaseStitch:
		return "INC"
	case DecreaseStitch:
		return "DEC"
	}

	panic("unreachable")
}

func getStitchKind(v string) StitchKind {
	switch v {
	case "k":
		return KnitStitch
	case "p":
		return PurlStitch
	case "s":
		return SlipStitch
	case "co", "c":
		return CastOnStitch
	case "bo", "b":
		return BindOffStitch
	case "inc":
		return IncreaseStitch
	case "dec":
		return DecreaseStitch
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	Kind StitchKind
}
