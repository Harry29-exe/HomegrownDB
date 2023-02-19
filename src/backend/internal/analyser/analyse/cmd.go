package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"log"
)

var CommandDelegator = commandDelegator{}

type commandDelegator struct{}

func (commandDelegator) Analyse(stmt pnode.CommandStmt, currentCtx anlctx.QueryCtx) (node.Query, error) {
	query := node.NewQuery(node.CommandTypeUtils, stmt)
	cmd := stmt.Stmt

	var commandStmt node.Node
	var err error
	switch cmd.Tag() {
	case pnode.TagCreateTable:
		commandStmt, err = CreateTable.Analyse(cmd.(pnode.CreateTableStmt), currentCtx)
	default:
		log.Panicf("not supported command type %d", cmd.Tag())
	}

	if err != nil {
		return nil, err
	}
	query.UtilsStmt = commandStmt
	return query, nil
}

// -------------------------
//      CreateTable
// -------------------------

var CreateTable = createTable{}

type createTable struct{}

func (createTable) Analyse(stmt pnode.CreateTableStmt, currentCtx anlctx.QueryCtx) (node.CreateRelation, error) {
	table := tabdef.NewTableDefinition(stmt.TableName)
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

	return node.NewCreateRelationTable(table), nil
}
