package pnode

import (
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

// InsertNode represents INSERT query
type InsertNode struct {
	Node
	Table    TableNode
	ColNames []string
	Rows     []InsertingRow
}

type InsertingRow struct {
	Node
	Values []Value
}

func NewInsertingRow() InsertingRow {
	return InsertingRow{
		Values: make([]Value, 0, 25),
	}
}

func (v *InsertingRow) AddValue(tk token.Token, tokenIndex uint32) bool {
	value := Value{Node: NewNode(tokenIndex, tokenIndex+1)}

	switch tk.Code() {
	case token.SqlTextValue:
		sqlTextTk := tk.(*token.SqlTextValueToken)
		value.V = sqlTextTk.InnerStr
	case token.Integer:
		value.V = tk.(*token.IntegerToken).Int
	case token.Float:
		value.V = tk.(*token.FloatToken).Float
	default:
		return false
	}

	v.Values = append(v.Values, value)
	return true
}
