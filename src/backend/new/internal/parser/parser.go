package parser

import (
	"HomegrownDB/backend/new/internal/parser/segparser"
	"HomegrownDB/backend/new/internal/parser/tokenizer"
	"HomegrownDB/backend/new/internal/parser/validator"
	"HomegrownDB/backend/new/internal/pnode"
)

func Parse(query string) (pnode.RawStmt, error) {
	src := tokenizer.NewTokenSource(query)
	v := validator.NewValidator(src)

	innerStmt, err := segparser.Delegator.Parse(src, v)
	if err != nil {
		return nil, err
	}

	stmt := pnode.NewRawStmt()
	stmt.Stmt = innerStmt
	return stmt, err
}
