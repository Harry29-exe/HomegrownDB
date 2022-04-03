package tokenizer

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

type SqlTextValueToken struct {
	code  TokenCode
	value string
	Str   string
}

func NewSqlTextValueToken(value string) *SqlTextValueToken {
	firstChar, lastChar := value[0], value[len(value)-1]
	if firstChar != '\'' || lastChar != '\'' {
		panic("given value can not be value of SqlTextValueToken because it does not have \"'\" signs on first and last position")
	}

	return &SqlTextValueToken{
		code:  SqlTextValue,
		value: value,
		Str:   value[1 : len(value)-1],
	}
}
