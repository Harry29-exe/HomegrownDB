package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/parser"
	tk "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

func Parse(query string) (Tree, error) {
	tokenSrc := internal.NewTokenSource(query)

	root, rootType, err := delegate(tokenSrc)
	if err != nil {
		return Tree{}, err
	}

	return Tree{
		RootType: rootType,
		Root:     root,
	}, nil
}

func delegate(source internal.TokenSource) (root any, rootType RootType, err error) {
	switch source.Current().Code() {
	case tk.Select:
		rootType = SELECT
		root, err = parser.Select.Parse(source)
	default:
		//todo implement me
		panic("Not implemented")
	}

	return
}
