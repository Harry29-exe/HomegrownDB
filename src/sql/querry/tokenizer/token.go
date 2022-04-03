package tokenizer

import "strconv"

type Token interface {
	Code() TokenCode
	Value() string
}

func NewBasicToken(code TokenCode, value string) Token {
	return &BasicToken{
		code:  code,
		value: value,
	}
}

type BasicToken struct {
	code  TokenCode
	value string
}

func (b *BasicToken) Code() TokenCode {
	return b.code
}

func (b *BasicToken) Value() string {
	return b.value
}

type TextToken struct {
	code            TokenCode
	value           string
	IsAscii         bool
	StartsWithDigit bool
}

func NewTextToken(value string) *TextToken {
	token := TextToken{
		code:    Text,
		value:   value,
		IsAscii: true,
	}

	for _, char := range value {
		if char > 127 {
			token.IsAscii = false
			break
		}
	}
	firstChar := value[0]
	if firstChar > 47 && firstChar < 58 {
		token.StartsWithDigit = true
	}

	return &token
}

func (t *TextToken) Code() TokenCode {
	return t.code
}

func (t *TextToken) Value() string {
	return t.value
}

func NewSqlTextValueToken(value string) (*SqlTextValueToken, error) {
	firstChar, lastChar := value[0], value[len(value)-1]
	if firstChar != '\'' || lastChar != '\'' {
		panic("given value can not be value of SqlTextValueToken because it does not have \"'\" signs on first and last position")
	}

	return &SqlTextValueToken{
		Token:  NewBasicToken(SqlTextValue, value),
		RawStr: value[1 : len(value)-1],
	}, nil
}

type SqlTextValueToken struct {
	Token
	RawStr string // RawStr is string inside quotation marks
}

func NewIntegerToken(value string) (*IntegerToken, error) {
	int, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return &IntegerToken{
		Token: NewBasicToken(Integer, value),
		Int:   int,
	}, nil
}

type IntegerToken struct {
	Token
	Int int
}

func NewFloatToken(value string) (*FloatToken, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	return &FloatToken{
		Token: NewBasicToken(Integer, value),
		Float: float,
	}, nil
}

type FloatToken struct {
	Token
	Float float64
}
