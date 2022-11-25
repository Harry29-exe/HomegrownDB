package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Parse(source internal.TokenSource, v validator.Validator) (pnode.InsertStmt, error) {
	source.Checkpoint()

	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		source.Rollback()
		return nil, err
	}
	source.Next()
	insertNode := pnode.NewInsertStmt()

	relation, err := Table.Parse(source, v)
	if err != nil {
		source.Rollback()
		return nil, err
	}
	insertNode.Relation = relation

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

	if err = v.CurrentIs(token.Values); err != nil {
		//todo implement me
		panic("Not implemented: insert with query is not yet supported")
	} else {
		insertValuesNode, err := InsertValues.Parse(source, v)
		if err != nil {
			source.Rollback()
			return insertNode, err
		}
		insertNode.Rows = insertValuesNode

		source.CommitAndInitNode(&insertNode.Node)
		return insertNode, nil
	}
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
