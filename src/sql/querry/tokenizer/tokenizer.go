package tokenizer

import (
	"errors"
	"unicode"
)

type Tokenizer interface {
	Next() (Token, error) //Next returns token, if there is no more tokens returns nil, if tokenizer encoders illegal char it returns error
	HasNext() bool
}

func NewTokenizer(str string) Tokenizer {
	return newBasicTokenizer(str)
}

type basicTokenizer struct {
	str         string //str string that tokenizer creates tokens from
	chars       []rune //chars str changed to rune slice
	len         uint   //length number of characters in str / len of chars
	pointer     uint   //pointer to next character to be parsed in chars
	futureToken []rune // futureToken slice pointing to chars describing next token value
}

func newBasicTokenizer(str string) *basicTokenizer {
	chars := []rune(str)
	charsLen := len(chars)
	return &basicTokenizer{
		str:         str,
		chars:       chars,
		pointer:     0,
		len:         uint(charsLen),
		futureToken: make([]rune, 0, charsLen/10),
	}
}

func isIn(char rune, chars []rune) bool {
	for i := 0; i < len(chars); i++ {
		if char == chars[i] {
			return true
		}
	}
	return false
}

func (t *basicTokenizer) Next() (Token, error) {
	err := t.createFutureToken()
	if err != nil {
		return nil, err
	}
	firstTokenChar := t.futureToken[0]
	switch {
	case isNonSpaceBreak(firstTokenChar):
		return t.tokenizeNonSpaceBreak()
	}
}

func (t *basicTokenizer) HasNext() bool {
	return t.pointer < t.len
}

func (t *basicTokenizer) tryParseNumber() {

}

func (t *basicTokenizer) tokenizeNonSpaceBreak() (Token, error) {
	switch t.futureToken[0] {
	case ',':
		return NewBasicToken(Comma, string(t.futureToken)), nil
	case '.':
		return NewBasicToken(Dot, string(t.futureToken)), nil
	case ';':
		return NewBasicToken(Semicolon, string(t.futureToken)), nil
	default:
		panic("unknown non-space character")
	}
}

func (t *basicTokenizer) tokenizeString() (Token, error) {
	switch t.futureToken[0] {
	case :
		
	}
}

func (t *basicTokenizer) createFutureToken() error {
	futureTokenStart := t.pointer

	t.pointer++
	nextChar := t.chars[t.pointer]

	for !isBreak(nextChar) {
		if unicode.IsControl(nextChar) {
			return errors.New("control character is not allowed in query")
		}

		t.pointer++
		nextChar = t.chars[t.pointer]
	}

	t.pointer++
	t.futureToken = t.chars[futureTokenStart:t.pointer]

	return nil
}
