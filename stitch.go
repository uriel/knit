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
)

func (s StitchKind) String() string {
	switch s {
	case KnitStitch:
		return "Knit"
	case PurlStitch:
		return "Purl"
	case SlipStitch:
		return "Slip"
	case CastOnStitch:
		return "Caston"
	case BindOffStitch:
		return "Bindoff"
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
	}

	return UnknownStitch
}

// A stich defines a specific kind of stitch to perform.
type Stitch struct {
	Kind StitchKind
}
