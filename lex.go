// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

import (
	"fmt"
	"io"
	"strings"
)

// lexState represents a single lexer state.
type lexState func(*lexer) lexState

// lexer is a lexer for knitting pattern strings.
type lexer struct {
	out      chan *token // Output channel for parsed tokens.
	state    lexState    // Current lexer state.
	data     string      // Input pattern string.
	line     [2]int      // Current line and line where token started.
	col      [2]int      // Current column and column where token started.
	lineSize int         // Size of previous line. Used for accurate rewind.
	start    int         // Start position of new token.
	pos      int         // Current end position of new token.
}

// lex reads the input data and turns it into a stream of tokens.
// tokens are sent over the returned channel.
func lex(data string) <-chan *token {
	l := new(lexer)

	if sz := len(data); sz == 0 {
		l.data = "\n"
	} else if data[sz-1] != '\n' {
		l.data = strings.ToLower(data) + "\n"
	}

	l.out = make(chan *token)
	l.state = lexText
	l.lineSize = 0
	l.line[0] = 1
	l.line[1] = 1
	l.col[0] = 1
	l.col[1] = 1
	l.start = 0
	l.pos = 0

	go func() {
		defer close(l.out)

		for l.state != nil {
			l.state = l.state(l)
		}
	}()

	return l.out
}

// emit emits an error token.
func (l *lexer) error(f string, argv ...interface{}) {
	l.out <- &token{tokError, fmt.Sprintf(f, argv...), l.line[1], l.col[1]}
	l.ignore()
}

// emit emits a new token.
func (l *lexer) emit(tt tokenType) {
	l.out <- &token{tt, l.data[l.start:l.pos], l.line[1], l.col[1]}
	l.ignore()
}

// current returns the currently read sequence of bytes.
func (l *lexer) current() string {
	return l.data[l.start:l.pos]
}

// next returns the next byte of data.
func (l *lexer) next() (byte, error) {
	if l.pos >= len(l.data) {
		l.emit(tokEof)
		return 0, io.EOF
	}

	b := l.data[l.pos]
	l.pos++

	if b == '\n' {
		l.line[0]++
		l.lineSize, l.col[0] = l.col[0], 0
	}

	l.col[0]++

	return b, nil
}

// rewind unreads the last byte.
func (l *lexer) rewind() {
	if l.pos > l.start {
		l.pos--
	}

	if l.col[0] > 1 {
		l.col[0]--
	} else {
		l.line[0]--
		l.col[0] = l.lineSize
	}
}

// ignore ignores any token data we have read so far.
func (l *lexer) ignore() {
	l.start = l.pos
	l.line[1] = l.line[0]
	l.col[1] = l.col[0]
}

// skip skips the next byte.
func (l *lexer) skip() {
	l.next()
	l.ignore()
}

// whitespace consumes bytes for as long as they qualify as whitespace.
func (l *lexer) whitespace() {
	for {
		b, err := l.next()

		if err != nil {
			return
		}

		if !isWhitespace(b) {
			l.rewind()
			break
		}
	}

	l.ignore()
	return
}

// number consumes bytes for as long as they qualify as digits.
func (l *lexer) number() bool {
	var n int

	for {
		b, err := l.next()

		if err != nil {
			return false
		}

		if !isDigit(b) {
			l.rewind()
			break
		}

		n++
	}

	if n > 0 {
		l.emit(tokNumber)
		return true
	}

	return false
}

// stitch consumes bytes for as long as they qualify as a known stitch name.
func (l *lexer) stitch() bool {
	var n int

	for {
		_, err := l.next()

		if err != nil {
			return false
		}

		if getStitchKind(l.current()) == UnknownStitch {
			l.rewind()
			break
		}

		n++
	}

	if n > 0 {
		l.emit(tokStitch)
		return true
	}

	return false
}

// reference consumes bytes for as long as they qualify as a reference.
// This means a name preceeded with a dollar sign '$'.
//
// A name is an ident that can additionally contain underscores and
// numbers.
func (l *lexer) reference() bool {
	b, err := l.next()

	if err != nil {
		return false
	}

	if b != '$' {
		l.rewind()
		return false
	}

	var n int
	for {
		b, err = l.next()

		if err != nil {
			return false
		}

		if !isName(b) {
			l.rewind()
			break
		}

		n++
	}

	if n > 0 {
		l.emit(tokReference)
		return true
	}

	return false
}

func isName(v byte) bool {
	return v == '_' || isDigit(v) || isLetter(v)
}

func isLetter(v byte) bool {
	return (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')
}

func isWhitespace(v byte) bool {
	switch v {
	case ' ', '\n', '\t', '\v', '\f', '\r', 0x85, 0xA0:
		return true
	}
	return false
}

func isDigit(v byte) bool {
	return v >= '0' && v <= '9'
}
