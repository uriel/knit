// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

// A number holds a concrete amount of times to repeat the preceeding
// group or stitch.
//
// For example: `P3`.
//
// The number `3` tells us the P stitch should be repeated
// exactly three times.
type Number struct {
	Value int64
}
