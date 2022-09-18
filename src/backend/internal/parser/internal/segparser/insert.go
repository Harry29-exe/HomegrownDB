package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var InsertParser = insertParser{}

type insertParser struct{}

func (i insertParser) Parse(source internal.TokenSource) (pnode.InsertNode, error) {
	v := validator.NewValidator(source)
	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		return pnode.InsertNode{}, err
	}
	source.Next()
	insertNode := pnode.InsertNode{}

	table, err := Table.Parse(source, v)
	if err != nil {
		return pnode.InsertNode{}, err
	}
	insertNode.Table = table

	err = v.NextSequence(token.SpaceBreak, token.OpeningParenthesis)
	if err == nil {
		insertCols, err := InsertColParser.Parse(source, v)
		if err != nil {
			return insertNode, err
		}
		insertNode.Columns.ColNames = insertCols.ColNames
	}
	err = v.NextIs(token.SpaceBreak)
	if err != nil {
		return insertNode, err
	}

	source.Next()
	insertValues, err := InsertValues.Parse(source, v)
	if err != nil {
		return insertNode, err
	}
	insertNode.Rows = insertValues

	return insertNode, nil
}
