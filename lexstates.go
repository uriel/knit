// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package knit

func lexText(l *lexer) lexState {
	l.whitespace()

	if l.literal("row") {
		l.emit(tokRow)
		return lexText
	}

	if l.number() {
		return lexText
	}

	if l.reference() {
		return lexText
	}

	if l.ident() {
		return lexText
	}

	c, _ := l.next()

	switch c {
	case '[':
		l.emit(tokGroupStart)
		return lexText
	case ']':
		l.emit(tokGroupEnd)
		return lexText

	// Punctuation sometimes used by users.
	// Don't consider it an error, just ignore it.
	case ':', ',', '.', ';':
		l.ignore()
		return lexText
	}

	l.error("Unexpected character %q", c)
	return nil
}
