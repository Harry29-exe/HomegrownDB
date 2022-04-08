package token

import (
	"errors"
	"strings"
)

type Code = uint16

const (
	/* -- KEYWORDS --*/

	Select Code = iota
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

	/* -- break characters -- */

	SpaceBreak
	Comma
	Dot
	Semicolon

	/* -- Values like string, int, float -- */

	Integer
	Float
	SqlTextValue

	/* -- Other -- */

	Text
	Nil
)

var breakCodes = map[Code]bool{
	SpaceBreak: true,
	Comma:      true,
	Dot:        true,
	Semicolon:  true,
}

func KeywordToToken(keyword string) (Token, error) {
	upperKeyword := strings.ToUpper(keyword)
	token, ok := keywordMap[upperKeyword]
	if !ok {
		return nil, errors.New("\"" + keyword + "\" is unknown keyword")
	}

	return NewBasicToken(token, keyword), nil
}

var keywordMap = map[string]Code{
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
