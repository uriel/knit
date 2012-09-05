// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"testing"
)

func TestPattern(t *testing.T) {
	tests := []string{
		`p k`, `p1 k1`,
		`co9 [p3 s3 k3]*10 bo`,
	}

	for i, str := range tests {
		_, err := Parse(str)

		if err != nil {
			t.Fatalf("Pattern %d: %v", i, err)
		}
	}
}
