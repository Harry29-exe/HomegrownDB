package pnode

import (
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"errors"
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
	Fields []InsertingField
}

func NewInsertingRow() InsertingRow {
	return InsertingRow{
		Fields: make([]InsertingField, 0, 25),
	}
}

type InsertingField struct {
	Select *Select
	Func   *rune // change rune it to Func pnode when it's implemented
	Value  *Value
}

func (v *InsertingRow) AppendValue(tk token.Token, tokenIndex uint32) error {
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
		return errors.New("token is not value")
	}

	v.Fields = append(v.Fields, InsertingField{Value: &value})
	return nil
}
