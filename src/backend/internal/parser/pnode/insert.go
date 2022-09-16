package pnode

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/sqlerr"
	"fmt"
)

// InsertNode represents INSERT query
type InsertNode struct {
	Table   TableNode
	Columns InsertingColumns
	Rows    []InsertingValues
}

type InsertingColumns struct {
	ColNames []string
}

type InsertingValues struct {
	Values []Value
}

func NewInsertingValue() InsertingValues {
	return InsertingValues{
		Values: make([]Value, 0, 25),
	}
}

func (v *InsertingValues) AddValue(tk token.Token, source internal.TokenSource) error {
	var value Value
	switch tk.Code() {
	case token.SqlTextValue:
		strTk := tk.(*token.SqlTextValueToken)
		value.V = strTk.RawStr
		value.Type = ValueTypeStr
	case token.Integer:
		intTk := tk.(*token.IntegerToken)
		value.V = intTk.Int
		value.Type = ValueTypeInt
	case token.Float:
		floatTk := tk.(*token.FloatToken)
		value.V = floatTk.Float
		value.Type = ValueTypeFloat
	default:
		return sqlerr.NewSyntaxError(
			"value that can be used as column value",
			fmt.Sprintf("got %s", token.ToString(tk.Code())),
			source,
		)
	}

	v.Values = append(v.Values, value)
	return nil
}
