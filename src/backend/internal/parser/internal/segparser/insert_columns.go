package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var InsertColParser = insertColParser{}

type insertColParser struct{}

// Parse when starting source should be at opening parenthesis position.
// Insert Into here -->(col1, col2, col3)
// Last token parsed is closing parenthesis
func (p insertColParser) Parse(source internal.TokenSource, v validator.Validator) (pnode.InsertingColumns, error) {
	err := v.CurrentIs(token.OpeningParenthesis)
	if err != nil {
		return pnode.InsertingColumns{}, err
	}

	err = v.SkipTokens().Type(token.SpaceBreak).SkipFromNext()
	if err != nil {
		return pnode.InsertingColumns{}, err
	}

	insertCols := pnode.InsertingColumns{ColNames: make([]string, 0, 10)}
	for {
		colName, err := v.Next().IsTextToken().AsciiOnly().DontStartWithDigit().Check()
		if err != nil {
			return pnode.InsertingColumns{}, err
		}
		insertCols.ColNames = append(insertCols.ColNames, colName.Value())

		err = v.SkipTokens().Type(token.SpaceBreak).TypeExactly(token.Comma, 1).SkipFromNext()
		if err != nil {
			_ = v.SkipTokens().Type(token.SpaceBreak).SkipFromNext()
			err = v.SkipTokens().
				TypeExactly(token.ClosingParenthesis, 1).
				SkipFromNext()

			return insertCols, err
		}
	}
}
