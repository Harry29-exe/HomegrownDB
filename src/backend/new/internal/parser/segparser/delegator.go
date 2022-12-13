package segparser

import (
	"HomegrownDB/backend/new/internal/parser/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
)

var Delegator = delegator{}

type delegator struct{}

func (delegator) Parse(src tkSource, v tkValidator) (pnode.Node, error) {
	switch tk := src.Current(); tk.Code() {
	case token.Select:
		return Select.Parse(src, v)
	case token.Insert:
		return Insert.Parse(src, v)
	default:
		panic("unsupported type: " + token.ToString(tk.Code()))
	}
}
