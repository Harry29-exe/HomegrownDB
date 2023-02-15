package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/pnode"
)

var CommandDelegator = commandDelegator{}

type commandDelegator struct{}

func (commandDelegator) Analyse(stmt pnode.CommandStmt, currentCtx anlsr.QueryCtx) error {
	//cmd := stmt.Stmt

	//todo implement me
	panic("Not implemented")
	//switch cmd.Tag() {
	//case pnode.TagCreateTable:
	//
	//}

}

// -------------------------
//      CreateTable
// -------------------------

var CreateTable = createTable{}

type createTable struct{}

func (createTable) Analyse(stmt pnode.CreateTableStmt, currentCtx anlsr.QueryCtx) error {
	//todo implement me
	panic("Not implemented")
}
