package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	tabdef "HomegrownDB/dbsystem/reldef/tabdef"
)

var CommandDelegator = commandDelegator{}

type commandDelegator struct{}

func (commandDelegator) Analyse(stmt pnode.CommandStmt, currentCtx anlctx.QueryCtx) error {
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

func (createTable) Analyse(stmt pnode.CreateTableStmt, currentCtx anlctx.QueryCtx) (node.CreateTable, error) {
	table := tabdef.NewDefinition(stmt.TableName)
	for i := 0; i < len(stmt.Columns); i++ {
		column, err := ColumnDef.Analyse(stmt.Columns[i], currentCtx)
		if err != nil {
			return nil, err
		}

		err = table.AddColumn(column)
		if err != nil {
			return nil, err
		}
	}

	return node.NewCreateTable(table), nil
}
