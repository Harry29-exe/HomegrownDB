package token

import (
	"errors"
	"strings"
)

func KeywordToToken(keyword string) (Token, error) {
	upperKeyword := strings.ToUpper(keyword)
	token, ok := keywordMap[upperKeyword]
	if !ok {
		return nil, errors.New("\"" + keyword + "\" is unknown keyword")
	}

	return NewBasicToken(token, keyword), nil
}

type Code = uint16

const (
	/* -- KEYWORDS --*/

	Select Code = iota
	From
	Where
	Join
	Update
	Insert
	Into
	Values
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
	OpeningParenthesis //todo add to tokenizer
	ClosingParenthesis //todo add to tokenizer

	/* -- Values like string, int, float -- */

	Integer
	Float
	SqlTextValue

	/* -- Other -- */

	Identifier
	Nil
	Error
)

var breakCodes = map[Code]bool{
	SpaceBreak: true,
	Comma:      true,
	Dot:        true,
	Semicolon:  true,
}

var keywordMap = map[string]Code{
	"SELECT": Select,
	"FROM":   From,
	"WHERE":  Where,
	"JOIN":   Join,
	"UPDATE": Update,
	"INSERT": Insert,
	"INTO":   Into,
	"VALUES": Values,
	"DELETE": Delete,
	"CREATE": Create,
	"DROP":   Drop,
	"TABLE":  Table,
	"AND":    And,
	"OR":     Or,
	"On":     On,
}
