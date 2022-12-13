package plan

import (
	"HomegrownDB/backend/internal/shared/qctx"
)

type ReduceFields struct {
	Fields []qctx.QColumnId
	Child  Node
}

func (s ReduceFields) Type() NodeType {
	return ReduceFieldsNode
}

func (s ReduceFields) Children() []Node {
	return []Node{s.Child}
}
