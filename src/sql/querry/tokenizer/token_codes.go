package tokenizer

import (
	"errors"
	"strings"
)

type TokenCode = uint16

const (
	/* -- KEYWORDS --*/

	Select TokenCode = iota
	From
	Where
	Join
	Update
	Delete
	Create
	Drop
	Table
	And
	Or
	On

	/* -- Non space break characters -- */

	Comma
	Dot
	Semicolon

	/* -- Values like string, int, float -- */

	Integer
	Float
	SqlTextValue

	/* -- Other -- */

	Text
)

func KeywordToToken(keyword string) (Token, error) {
	upperKeyword := strings.ToUpper(keyword)
	token, ok := keywordMap[upperKeyword]
	if !ok {
		return nil, errors.New("\"" + keyword + "\" is unknown keyword")
	}

	return NewBasicToken(token, keyword), nil
}

var keywordMap = map[string]TokenCode{
	"SELECT": Select,
	"FROM":   From,
	"WHERE":  Where,
	"JOIN":   Join,
	"UPDATE": Update,
	"DELETE": Delete,
	"CREATE": Create,
	"DROP":   Drop,
	"TABLE":  Table,
	"AND":    And,
	"OR":     Or,
	"On":     On,
}
