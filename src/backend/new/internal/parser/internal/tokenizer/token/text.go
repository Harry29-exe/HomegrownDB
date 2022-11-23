package token

func NewTextToken(value string) *TextToken {
	token := TextToken{
		code:    Identifier,
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

type TextToken struct {
	code            Code
	value           string
	IsAscii         bool
	StartsWithDigit bool
}

func (t *TextToken) Code() Code {
	return t.code
}

func (t *TextToken) Value() string {
	return t.value
}
