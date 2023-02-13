package parser

import (
	"HomegrownDB/backend/internal/parser/parse"
	"HomegrownDB/backend/internal/parser/tokenizer"
	"HomegrownDB/backend/internal/parser/validator"
	"HomegrownDB/backend/internal/pnode"
)

func Parse(query string) (pnode.RawStmt, error) {
	src := tokenizer.NewTokenSource(query)
	v := validator.NewValidator(src)

	innerStmt, err := parse.Delegator.Parse(src, v)
	if err != nil {
		return nil, err
	}

	stmt := pnode.NewRawStmt()
	stmt.Stmt = innerStmt
	return stmt, err
}
