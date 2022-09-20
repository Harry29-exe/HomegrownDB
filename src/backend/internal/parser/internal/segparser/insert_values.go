package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
	"fmt"
)

var InsertValues = insertValues{}

type insertValues struct{}

// Parse part of query starting with Values or Value token.
// String to parse by this method can look like this:
//
// VALUES (colName1, colName2 , colName4,colName6), ( colName1,colName2 ,colName4 , colName6 )
//
// It will start parsing it when current token is Values and return
// when source pointer is on closing parenthesis
func (i insertValues) Parse(source internal.TokenSource, v validator.Validator) ([]pnode.InsertingRow, error) {
	source.Checkpoint()
	err := v.CurrentIsAnd(token.Values).
		SkipTokens().
		TypeMax(token.SpaceBreak, 1).
		SkipFromNext()
	if err != nil {
		return nil, err
	}

	values := make([]pnode.InsertingRow, 0, 25)
	var value pnode.InsertingRow

	source.Next()
	for {
		value, err = i.parseRow(source, v)
		if err != nil {
			source.Rollback()
			return values, err
		}

		values = append(values, value)

		err = v.SkipTokens().
			TypeExactly(token.Comma, 1).
			TypeMax(token.SpaceBreak, 2).
			SkipFromNext()
		if err != nil {
			if len(values) > 0 {
				source.Rollback()
				return values, nil
			}
			source.Commit()
			return values, err
		}

		err = v.NextIs(token.OpeningParenthesis)
		if err != nil {
			source.Rollback()
			return values, err
		}
	}
}

func (i insertValues) parseRow(source internal.TokenSource, v validator.Validator) (insertingRow pnode.InsertingRow, err error) {
	if v.CurrentIs(token.OpeningParenthesis) != nil {
		return
	} else {
		_ = v.SkipSpaceFromNext()
	}

	rowNode := pnode.NewInsertingRow()
	if !rowNode.AddValue(source.Next()) {
		return pnode.InsertingRow{}, sqlerr.NewSyntaxError(
			"value that can be used as column value",
			fmt.Sprintf("got %s", token.ToString(source.Current().Code())),
			source,
		)
	}

	for {
		if v.SkipTokens().
			TypeExactly(token.Comma, 1).
			TypeMax(token.SpaceBreak, 2).
			SkipFromNext() != nil {

			_ = v.SkipSpaceFromNext()
			nextTk := source.Next()
			if nextTk.Code() != token.ClosingParenthesis {
				return rowNode, sqlerr.NewSyntaxError(")", token.ToString(nextTk.Code()), source)
			}

			return rowNode, nil

		} else if !rowNode.AddValue(source.Next()) {
			return pnode.InsertingRow{}, sqlerr.NewSyntaxError(
				"value that can be used as column value",
				fmt.Sprintf("got %s", token.ToString(source.Current().Code())),
				source,
			)
		}
	}
}
