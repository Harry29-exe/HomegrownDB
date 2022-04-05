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

func ToString(code TokenCode) string {
	str, ok := tokenNamesMap[code]
	if ok {
		return str
	}
	return ""
}

const (
	SelectStr = "SELECT"
	FromStr   = "FROM"
	WhereStr  = "WHERE"
	JoinStr   = "JOIN"
	UpdateStr = "UPDATE"
	DeleteStr = "DELETE"
	CreateStr = "CREATE"
	DropStr   = "DROP"
	TableStr  = "TABLE"
	AndStr    = "AND"
	OrStr     = "OR"
	OnStr     = "ON"

	/* -- break characters -- */

	SpaceBreakStr = "SpaceBreak"
	CommaStr      = "Comma"
	DotStr        = "Dot"
	SemicolonStr  = "Semicolon"

	/* -- Values like string, int, float -- */

	IntegerStr      = "Integer"
	FloatStr        = "Float"
	SqlTextValueStr = "SqlTextValue"

	/* -- Other -- */

	TextStr = "Text"
)

var tokenNamesMap = map[TokenCode]string{
	Select: SelectStr,
	From:   FromStr,
	Where:  WhereStr,
	Join:   JoinStr,
	Update: UpdateStr,
	Delete: DeleteStr,
	Create: CreateStr,
	Drop:   DropStr,
	Table:  TableStr,
	And:    AndStr,
	Or:     OrStr,
	On:     OnStr,

	/* -- break characters -- */

	SpaceBreak: SpaceBreakStr,
	Comma:      CommaStr,
	Dot:        DotStr,
	Semicolon:  SemicolonStr,

	/* -- Values like string, int, float -- */

	Integer:      IntegerStr,
	Float:        FloatStr,
	SqlTextValue: SqlTextValueStr,

	/* -- Other -- */

	Text: TextStr,
}
