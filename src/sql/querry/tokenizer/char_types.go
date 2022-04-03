package tokenizer

import "unicode"

var breakChars = newCharCollection([]rune{',', '.', ';'})

func isStrAppendable(char rune) bool {
	switch {
	case breakChars.In(char) || unicode.IsSpace(char):
		return false
	case unicode.IsControl(char):
		panic("Control char in query string")
	default:
		return true
	}
}

func isNonSpaceBreak(char rune) bool {
	return breakChars.In(char)
}

func isBreak(char rune) bool {
	return unicode.IsSpace(char) || isNonSpaceBreak(char)
}
