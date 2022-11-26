package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
)

var Values = values{}

type values struct{}

func (val values) Parse(src tkSource, v tkValidator) ([]pnode.Node, error) {
	nodes := make([]pnode.Node, 0, 10)

	node, err := Value.Parse(src, v)
	if err != nil {
		return nil, err
	}
	nodes = append(nodes, node)

	for val.hasNextVal(v) {
		node, err = Value.Parse(src, v)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (val values) hasNextVal(v tkValidator) bool {
	return v.SkipTokens().
		Type(token.Comma).
		TypeMax(token.SpaceBreak, 2).
		SkipFromCurrent() == nil
}

var Value = value{}

type value struct{}

func (val value) Parse(src tkSource, v tkValidator) (pnode.Node, error) {
	src.Checkpoint()
	var node pnode.Node
	var err error

	if val.isFunction(v) {
		src.Rollback()
		//todo implement me
		panic("Not implemented")
	} else if val.isColumnRef(v) {
		src.Rollback()
		node, err = val.parseColumnRef(src, v)
	} else {
		panic("not supported")
	}
	//todo add support for expressions

	return node, err
}

func (val value) isFunction(v tkValidator) bool {
	return v.CurrentSequence(token.Identifier, token.OpeningParenthesis) == nil
}

func (val value) isColumnRef(v tkValidator) bool {
	return v.CurrentSequence(token.Identifier, token.SpaceBreak) == nil ||
		v.CurrentSequence(token.Identifier, token.Dot, token.Identifier) == nil
}

func (val value) parseColumnRef(src tkSource, v tkValidator) (pnode.ColumnRef, error) {
	identifier1 := src.Current()
	if tk := src.Next(); tk.Code() == token.SpaceBreak {
		return pnode.NewColumnRef(identifier1.Value(), ""), nil
	} else {
		cRef := pnode.NewColumnRef(src.Next().Value(), identifier1.Value())
		src.Next()
		return cRef, nil
	}
}
