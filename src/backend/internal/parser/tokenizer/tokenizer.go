package tokenizer

import (
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	"errors"
	"unicode"
)

type Tokenizer interface {
	Next() (token2.Token, error) //Next returns token, if there is no more tokens returns nil, if tokenizer encoders illegal char it returns error
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

func (t *basicTokenizer) Next() (token2.Token, error) {
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

func (t *basicTokenizer) tokenizeNumber() (tk token2.Token, err error) {
	for _, char := range t.futureToken {
		if char == '.' {
			tk, err = token2.NewFloatToken(string(t.futureToken))
			return
		}
	}
	tk, err = token2.NewIntegerToken(string(t.futureToken))
	return
}

func (t *basicTokenizer) tokenizeSpaceBreak() (token2.Token, error) {
	return token2.NewBasicToken(token2.SpaceBreak, " "), nil
}

func (t *basicTokenizer) tokenizeNonSpaceBreak() (token2.Token, error) {
	switch t.futureToken[0] {
	case ',':
		return token2.NewBasicToken(token2.Comma, string(t.futureToken)), nil
	case '.':
		return token2.NewBasicToken(token2.Dot, string(t.futureToken)), nil
	case ';':
		return token2.NewBasicToken(token2.Semicolon, string(t.futureToken)), nil
	case ')':
		return token2.NewBasicToken(token2.ClosingParenthesis, string(t.futureToken)), nil
	case '(':
		return token2.NewBasicToken(token2.OpeningParenthesis, string(t.futureToken)), nil
	default:
		panic("unknown non-space character")
	}
}

func (t *basicTokenizer) tokenizeString() (token2.Token, error) {
	keywordToken, err := token2.KeywordToToken(string(t.futureToken))
	if err == nil {
		return keywordToken, nil
	}

	return token2.NewTextToken(string(t.futureToken)), nil
}

func (t *basicTokenizer) tokenizeSqlString() (token2.Token, error) {
	return token2.NewSqlTextValueToken(string(t.futureToken))
}

// todo add support for sql text eg. 'some text'
func (t *basicTokenizer) createFutureToken() error {
	if t.pointer >= t.len {
		return errors.New("tokenizer has no more tokens")
	}

	futureTokenStart := t.pointer
	nextChar := t.chars[t.pointer]
	switch {
	case unicode.IsSpace(nextChar):
		t.pointer++
		t.futureToken = t.chars[futureTokenStart:t.pointer]
		t.skipWhiteSpaces()
	case nextChar == '\'':
		t.pointer++
		if t.pointer >= t.len {
			return eofError(token2.ToString(token2.SqlTextValue), t.str)
		}
		nextChar = t.chars[t.pointer]
		for nextChar != '\'' {
			t.pointer++
			if t.pointer >= t.len {
				return eofError(token2.ToString(token2.SqlTextValue), t.str)
			}
			nextChar = t.chars[t.pointer]
		}
		t.pointer++
	default:
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
