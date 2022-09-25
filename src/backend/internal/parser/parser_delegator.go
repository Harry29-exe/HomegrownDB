package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/segparser"
	tk "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/dbsystem/tx"
)

func Parse(query string, ctx *tx.Ctx) (Tree, error) {
	tokenSrc := internal.NewTokenSource(query)

	root, rootType, err := delegate(tokenSrc)
	if err != nil {
		return Tree{}, err
	}

	setTokenHistory(ctx, tokenSrc)
	return Tree{
			RootType: rootType,
			Root:     root,
		},
		nil
}

func setTokenHistory(ctx *tx.Ctx, source internal.TokenSource) {
	tokens := source.History()
	values := make([]string, len(tokens))
	for i := 0; i < len(tokens); i++ {
		values[i] = tokens[i].Value()
	}

	ctx.CurrentQuery.QueryTokens = values
}

func delegate(source internal.TokenSource) (root any, rootType RootType, err error) {
	v := validator.NewValidator(source)
	switch source.Current().Code() {
	case tk.Select:
		rootType = Select
		root, err = segparser.Select.Parse(source, v)
	default:
		rootType = Insert
		root, err = segparser.Insert.Parse(source, v)
	}

	return
}
