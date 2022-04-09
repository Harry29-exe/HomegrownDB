package tokenizer

import (
	. "HomegrownDB/sql/backend/tokenizer/token"
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
	bt := &basicTokenizer{
		str:         str,
		chars:       chars,
		pointer:     0,
		len:         uint(charsLen),
		futureToken: make([]rune, 0, charsLen/10),
	}

	bt.skipWhiteSpaces()
	return bt
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
	case unicode.IsSpace(firstTokenChar):
		return t.tokenizeSpaceBreak()
	case isNonSpaceBreak(firstTokenChar):
		return t.tokenizeNonSpaceBreak()
	case unicode.IsDigit(firstTokenChar):
		return t.tokenizeNumber()
	case firstTokenChar == '\'':
		return t.tokenizeSqlString()
	case unicode.IsDigit(firstTokenChar):
		token, err := t.tokenizeNumber()
		if err != nil {
			return t.tokenizeString()
		}
		return token, nil
	default:
		return t.tokenizeString()
	}
}

func (t *basicTokenizer) HasNext() bool {
	return t.pointer < t.len
}

func (t *basicTokenizer) tokenizeNumber() (token Token, err error) {
	for _, char := range t.futureToken {
		if char == '.' {
			token, err = NewFloatToken(string(t.futureToken))
			return
		}
	}
	token, err = NewIntegerToken(string(t.futureToken))
	return
}

func (t *basicTokenizer) tokenizeSpaceBreak() (Token, error) {
	return NewBasicToken(SpaceBreak, " "), nil
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
	keywordToken, err := KeywordToToken(string(t.futureToken))
	if err == nil {
		return keywordToken, nil
	}

	return NewTextToken(string(t.futureToken)), nil
}

func (t *basicTokenizer) tokenizeSqlString() (Token, error) {
	return NewSqlTextValueToken(string(t.futureToken))
}

//todo add support for sql text eg. 'some text'
func (t *basicTokenizer) createFutureToken() error {
	if t.pointer >= t.len {
		return errors.New("tokenizer has no more tokens")
	}

	futureTokenStart := t.pointer
	nextChar := t.chars[t.pointer]
	if unicode.IsSpace(nextChar) {
		t.pointer++
		t.futureToken = t.chars[futureTokenStart:t.pointer]
		t.skipWhiteSpaces()
	}

	for !isBreak(nextChar) {
		if unicode.IsControl(nextChar) {
			return errors.New("control character is not allowed in query")
		}

		t.pointer++
		if t.pointer >= t.len {
			break
		}
		nextChar = t.chars[t.pointer]
	}

	if t.pointer == futureTokenStart {
		t.pointer++
	}
	t.futureToken = t.chars[futureTokenStart:t.pointer]

	//t.skipWhiteSpaces()
	return nil
}

func (t *basicTokenizer) skipWhiteSpaces() {
	if t.pointer >= t.len {
		return
	}

	char := t.chars[t.pointer]
	for unicode.IsSpace(char) {
		t.pointer++
		if t.pointer >= t.len {
			break
		}

		char = t.chars[t.pointer]
	}
}
