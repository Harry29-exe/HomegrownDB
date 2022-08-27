package parser

import (
	"HomegrownDB/backend/parser/internal/parsers"
	"HomegrownDB/backend/parser/internal/source"
	tk "HomegrownDB/backend/parser/tokenizer/token"
)

func Parse(query string) (Tree, error) {
	tokenSrc := NewTokenSource(query)

	root, rootType, err := delegate(tokenSrc)
	if err != nil {
		return nil, err
	}

	return BasicTree{
		rootType: rootType,
		root:     root,
	}, nil
}

func delegate(source source.TokenSource) (root any, rootType RootType, err error) {
	switch source.Current().Code() {
	case tk.Select:
		rootType = SELECT
		root, err = parsers.Select.Parse(source)
	default:
		//todo implement me
		panic("Not implemented")
	}

	return
}
