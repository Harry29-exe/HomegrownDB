package parse

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
)

var Command = command{}

type command struct{}

func (c command) Parse(src tkSource, v tkValidator) (pnode.CommandStmt, error) {
	src.Checkpoint()
	var stmt pnode.Node
	var err error

	firstTk := src.Current()
	if firstTk.Code() == token.Create {
		if err = v.NextIs(token.SpaceBreak); err != nil {
			return nil, err
		}
		secondTk := src.Next()
		switch secondTk.Code() {
		case token.Table:
			src.Rollback()
			stmt, err = CreateTable.Parse(src, v)
			if err != nil {
				return nil, err
			}
		default:
			return nil, sqlerr.NewSyntaxError("TABLE|?|?", token.ToString(secondTk.Code()), src)
		}
	} else {
		//todo implement me
		panic("Not implemented")
	}

	return pnode.NewCommandStmt(stmt), nil
}

// -------------------------
//      CreateTableStmt
// -------------------------

var CreateTable = createTable{}

type createTable struct{}

func (c createTable) Parse(src tkSource, v tkValidator) (pnode.CreateTableStmt, error) {
	err := v.CurrentSequence(token.Create, token.SpaceBreak, token.Table, token.SpaceBreak, token.Identifier)
	if err != nil {
		return nil, err
	}
	tableName := src.Current().Value()
	err = v.NextSequence(token.SpaceBreak, token.OpeningParenthesis, token.SpaceBreak)
	if err != nil {
		return nil, err
	}
	src.Next()

	columns := make([]pnode.ColumnDef, 1, 10)
	if columns[0], err = ColumnDef.Parse(src, v); err != nil {
		return nil, err
	}
	for c.nextColumn(src, v) {
		col, err := ColumnDef.Parse(src, v)
		if err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}

	return pnode.NewCreateTableStmt(tableName, columns), nil
}

func (c createTable) nextColumn(src tkSource, v tkValidator) (success bool) {
	src.Checkpoint()

	_ = v.SkipCurrentSB()
	tk := src.Current()
	if tk.Code() != token.Comma {
		src.Rollback()
		return false
	}
	_ = v.SkipNextSB()
	src.Commit()
	return true
}
