package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var InsertParser = insertParser{}

type insertParser struct{}

func (i insertParser) Parse(source internal.TokenSource) error {
	v := validator.NewValidator(source)
	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		return err
	}
	source.Next()
	insertNode := pnode.InsertNode{}

	table, err := Table.Parse(source, v)
	if err != nil {
		return err
	}
	insertNode.Table = table

	err = v.NextSequence(token.SpaceBreak, token.OpeningParenthesis)
	if err != nil {

	}

	//todo implement me
	panic("Not implemented")
}
