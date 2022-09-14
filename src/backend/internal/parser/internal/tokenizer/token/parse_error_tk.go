package token

func NewErrorToken(errorMsg string) Token {
	return ErrorToken{tokenizerErrorMsg: errorMsg}
}

var (
	_ error = ErrorToken{}
	_ Token = ErrorToken{}
)

type ErrorToken struct {
	tokenizerErrorMsg string
}

func (e ErrorToken) Error() string {
	return e.tokenizerErrorMsg
}

func (e ErrorToken) Code() Code {
	return Error
}

func (e ErrorToken) Value() string {
	return ""
}
