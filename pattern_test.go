// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"os"
	"testing"
)

func TestPattern(t *testing.T) {
	tests := []string{
		`p k`, `p1 k1`,
		`c9
		[p3 yo s3 yo k3 inc]10
		bo+`,
	}

	for i, str := range tests {
		p, err := Parse(fmt.Sprintf("Pattern %d", i), str)

		if err != nil {
			t.Fatal(err)
		}

		p.dump(os.Stdout)
	}
}
