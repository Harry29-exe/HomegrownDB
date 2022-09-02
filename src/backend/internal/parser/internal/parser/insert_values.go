package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
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
func (i insertValsParser) Parse(source internal.TokenSource, v validator.Validator) ([]pnode.InsertingValues, error) {
	err := v.CurrentIsAnd(token.Values).
		SkipBreaks().
		TypeMax(token.SpaceBreak, 1).
		SkipFromNext()
	if err != nil {
		return nil, err
	}

	values := make([]pnode.InsertingValues, 0, 25)
	var value pnode.InsertingValues

	source.Next()
	for {
		value, err = i.parseValue(source, v)
		if err != nil {
			return values, err
		}

		values = append(values, value)

		err = v.SkipBreaks().
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

func (i insertValsParser) parseValue(source internal.TokenSource, v validator.Validator) (pnode.InsertingValues, error) {
	//todo implement me
	panic("Not implemented")
}
