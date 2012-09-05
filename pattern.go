// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
)

// Pattern represents a single, complete knitting pattern.
type Pattern struct {
}

// MustParse parses the input pattern.
// It panics if an error occurred.
func MustParse(pat string) *Pattern {
	p, err := Parse(pat)

	if err != nil {
		panic(err)
	}

	return p
}

// Parse parses the given input pattern.
func Parse(pat string) (*Pattern, error) {
	p := new(Pattern)
	tokens := lex(pat)

loop:
	for {
		select {
		case tok := <-tokens:
			if tok == nil || tok.Type == tokEof {
				break loop
			}

			switch tok.Type {
			case tokError:
				return nil, fmt.Errorf("%d:%d %s",
					tok.Line, tok.Col, tok.Data)

			}

			//fmt.Printf("%6s: %q\n", tok.Type, tok.Data)
		}
	}

	return p, nil
}
