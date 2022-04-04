package leaf

import (
	"HomegrownDB/sql/querry/parser/common"
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
)

var FieldParser fieldParser = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source common.TokenSource) (parsetree.Node, error) {
	source.Checkpoint()
	tableToken := source.Next()
	if tableToken.Code() != tokenizer.Text {
		defer source.Rollback()
		return nil, common.NewSyntaxError("table alias", tableToken.Value(), source)
	}
	tableAliasToken := tableToken.(*tokenizer.TextToken)
	//todo move this checks to some shared function
	if tableAliasToken.StartsWithDigit {
		defer source.Rollback()
		return nil, common.NewSyntaxTextError("table alias can not start with digit", source)
	} else if !tableAliasToken.IsAscii {
		defer source.Rollback()
		return nil, common.NewSyntaxTextError("table alias must contain only ascii chars", source)
	}

	dot := source.Next()
	if dot.Code() != tokenizer.Dot {
		defer source.Rollback()
		return nil, common.NewSyntaxError(".", dot.Value(), source)
	}

	//todo add check similar to table alias
	columnToken := source.Next()
	if columnToken.Code() != tokenizer.Text {
		defer source.Rollback()
		return nil, common.NewSyntaxError("column name", columnToken.Value(), source)
	}

	source.Commit()
	return parsetree.NewFieldNode(parsetree.FieldNodeValue{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: "",
	}), nil
}
