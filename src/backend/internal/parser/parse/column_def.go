package parse

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
)

var ColumnDef = columnDef{}

type columnDef struct{}

func (c columnDef) Parse(src tkSource, v tkValidator) (pnode.ColumnDef, error) {
	columnNameTk, err := v.Current().
		Has(token.Identifier).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()
	if err != nil {
		return nil, err
	}
	columnName := columnNameTk.Value()

	err = v.NextIs(token.SpaceBreak)
	if err != nil {
		return nil, err
	}
	typeNameTk, err := v.Next().
		Has(token.Identifier).
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		return nil, err
	}
	typeName := typeNameTk.Value()
	_ = src.Next()
	_ = v.SkipCurrentSB()

	args, err := c.parseTypeConstrains(src, v)
	if err != nil {
		return nil, err
	}

	return pnode.NewColumnDef(columnName, typeName, args), nil
}

func (c columnDef) parseTypeConstrains(src tkSource, v tkValidator) ([]pnode.TypeArg, error) {
	src.Checkpoint()
	args := make([]pnode.TypeArg, 0, 5)
	// checking for length
	if src.Current().Code() == token.OpeningParenthesis {
		lengthArg, err := c.parseLengthArgs(src, v)
		if err != nil {
			src.Rollback()
			return nil, err
		}

		args = append(args, lengthArg)
	}

	//for c.hasNextArg(src) {
	//
	//}
	//todo add support for more args
	src.Commit()
	return args, nil
}

func (c columnDef) parseLengthArgs(src tkSource, v tkValidator) (pnode.TypeArg, error) {
	_ = src.Next()
	_ = v.SkipCurrentSB()
	lenTk := src.Current()
	if lenTk.Code() != token.Integer {
		return nil, sqlerr.NewTokenSyntaxError(token.Integer, lenTk.Code(), src)
	}
	lenIntTk := lenTk.(*token.IntegerToken)
	arg := pnode.NewArgLength(int(lenIntTk.Int))

	_ = src.Next()
	if err := v.SkipCurrentSBAnd().CurrentIs(token.ClosingParenthesis); err != nil {
		return arg, err
	}
	return arg, nil
}

//func (c columnDef) parseArg(src tkSource, v tkValidator) pnode.TypeArg {
//	arg := src.Next()
//	switch arg.Value() {
//	case :
//
//	}
//}
//
//func (c columnDef) hasNextArg(src tkSource) bool {
//	src.Checkpoint()
//	defer src.Rollback()
//	if src.Current().Code() != token.SpaceBreak {
//		return false
//	}
//	argTk := src.Next()
//	if argTk.Code() != token.Identifier && argTk.Code() != token.{
//		return false
//	}
//	return true
//}
