package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/parser/sqlerr"
)

var InsertValues = insertValsParser{}

type insertValsParser struct{}

// Parse part of query starting with Values or Value token.
// String to parse by this method can look like this:
//
// VALUES (colName1, colName2 , colName4,colName6), ( colName1,colName2 ,colName4 , colName6 )
//
// It will start parsing it when current token is Values and return
// when source pointer is on closing parenthesis
func (i insertValsParser) Parse(source internal.TokenSource, v validator.Validator) ([]pnode.InsertingRow, error) {
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
		value, err = i.parseValue(source, v)
		if err != nil {
			return values, err
		}

		values = append(values, value)

		err = v.SkipTokens().
			TypeExactly(token.Comma, 1).
			TypeMax(token.SpaceBreak, 2).
			SkipFromNext()
		if err != nil {
			if len(values) > 0 {
				return values, nil
			}
			return values, err
		}

		err = v.NextIs(token.OpeningParenthesis)
		if err != nil {
			return nil, err
		}
	}
}

func (i insertValsParser) parseValue(source internal.TokenSource, v validator.Validator) (pnode.InsertingRow, error) {
	err := v.CurrentIsAnd(token.OpeningParenthesis).
		SkipTokens().
		TypeMax(token.SpaceBreak, 1).
		SkipFromNext()
	if err != nil {
		return pnode.InsertingRow{}, err
	}

	values := pnode.NewInsertingValue()
	err = values.AddValue(source.Next(), source)
	if err != nil {
		return pnode.InsertingRow{}, err
	}

	for {
		err = v.SkipTokens().
			TypeExactly(token.Comma, 1).
			TypeMax(token.SpaceBreak, 2).
			SkipFromNext()
		if err != nil {
			nextTk := source.Next()
			if nextTk.Code() == token.SpaceBreak {
				nextTk = source.Next()
			}
			if nextTk.Code() != token.ClosingParenthesis {
				return values, sqlerr.NewSyntaxError(")", token.ToString(nextTk.Code()), source)
			}
			return values, nil
		}
		err = values.AddValue(source.Next(), source)
		if err != nil {
			return pnode.InsertingRow{}, err
		}
	}
}
