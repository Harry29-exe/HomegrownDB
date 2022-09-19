package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Parse(source internal.TokenSource, v validator.Validator) (pnode.InsertNode, error) {
	source.Checkpoint()
	startToken := source.CurrentTokenIndex()

	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		source.Rollback()
		return pnode.InsertNode{}, err
	}
	source.Next()
	insertNode := pnode.InsertNode{}

	table, err := Table.Parse(source, v)
	if err != nil {
		source.Rollback()
		return pnode.InsertNode{}, err
	}
	insertNode.Table = table

	err = v.NextSequence(token.SpaceBreak, token.OpeningParenthesis)
	if err == nil {
		err = i.parseInsertingCols(&insertNode, v)
		if err != nil {
			source.Rollback()
			return insertNode, err
		}
	}
	err = v.NextIs(token.SpaceBreak)
	if err != nil {
		source.Rollback()
		return insertNode, err
	}

	source.Next()
	insertValues, err := InsertValues.Parse(source, v)
	if err != nil {
		source.Rollback()
		return insertNode, err
	}
	insertNode.Rows = insertValues

	source.Commit()
	insertNode.SetTokenIndexes(startToken, source.CurrentTokenIndex())
	return insertNode, nil
}

func (i insert) parseInsertingCols(insertNode *pnode.InsertNode, v validator.Validator) error {
	err := v.CurrentIs(token.OpeningParenthesis)
	if err != nil {
		return err
	}

	err = v.SkipTokens().Type(token.SpaceBreak).SkipFromNext()
	if err != nil {
		return err
	}

	colNames := make([]string, 0, 10)
	var colName token.Token
	for {
		colName, err = v.Next().
			IsTextToken().
			AsciiOnly().
			DontStartWithDigit().
			Check()
		if err != nil {
			return err
		}
		colNames = append(colNames, colName.Value())

		err = v.SkipTokens().
			Type(token.SpaceBreak).
			TypeExactly(token.Comma, 1).
			SkipFromNext()

		if err != nil {
			insertNode.ColNames = colNames
			_ = v.SkipTokens().Type(token.SpaceBreak).SkipFromNext()

			return v.SkipTokens().
				TypeExactly(token.ClosingParenthesis, 1).
				SkipFromNext()
		}
	}
}
