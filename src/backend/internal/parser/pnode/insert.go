package pnode

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"fmt"
)

// InsertNode represents INSERT query
type InsertNode struct {
	Node
	Table    TableNode
	ColNames []string
	Rows     []InsertingRow
}

type InsertingRow struct {
	Values []Value
}

func NewInsertingValue() InsertingRow {
	return InsertingRow{
		Values: make([]Value, 0, 25),
	}
}

func (v *InsertingRow) AddValue(tk token.Token, source internal.TokenSource) error {
	var value Value
	switch tk.Code() {
	case token.SqlTextValue:
		value = StrValue{v: tk.(*token.SqlTextValueToken).InputStr}
	case token.Integer:
		value = IntValue{v: tk.(*token.IntegerToken).Int}
	case token.Float:
		value = FloatValue{v: tk.(*token.FloatToken).Float}
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
