package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Table = tableParser{}

type tableParser struct {
}

func (t tableParser) Parse(source internal.TokenSource, validator validator.Validator) (pnode.TableNode, error) {
	source.Checkpoint()

	name, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return pnode.TableNode{}, err
	}

	err = validator.NextSequence(token.SpaceBreak, token.Identifier)
	if err != nil {
		source.Commit()
		return pnode.TableNode{TableName: name.Value(), TableAlias: name.Value()}, nil
	}

	alias, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return pnode.TableNode{}, err
	}

	source.Commit()
	return pnode.TableNode{
		TableName:  name.Value(),
		TableAlias: alias.Value(),
	}, nil
}
