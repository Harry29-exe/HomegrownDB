package parse

import (
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
)

var Delegator = delegator{}

type delegator struct{}

func (delegator) Parse(src tkSource, v tkValidator) (pnode.Node, error) {
	switch tk := src.Current(); tk.Code() {
	case token2.Select:
		return Select.Parse(src, v)
	case token2.Insert:
		return Insert.Parse(src, v)
	case token2.Values:
		return ValueStreamSelect.Parse(src, v)
	default:
		return Command.Parse(src, v)
	}
}
