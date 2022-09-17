package token

func NewSqlTextValueToken(value string) (*SqlTextValueToken, error) {
	firstChar, lastChar := value[0], value[len(value)-1]
	if firstChar != '\'' || lastChar != '\'' {
		panic("given value can not be value of SqlTextValueToken because it does not have \"'\" signs on first and last position")
	}

	return &SqlTextValueToken{
		Token:    NewBasicToken(SqlTextValue, value),
		InputStr: value[1 : len(value)-1],
	}, nil
}

type SqlTextValueToken struct {
	Token
	InputStr string // InputStr is string inside quotation marks
}
