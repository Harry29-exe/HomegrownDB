package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

var Tables = tables{}

type tables struct{}

func (t tables) Analise(node []pnode.TableNode, ctx *tx.Ctx) (anode.Tables, error) {
	tablesCount := len(node)
	qTables := make([]anode.Table, tablesCount)

	var tableDef table.Definition
	var err error
	for i := 0; i < tablesCount; i++ {
		tableDef, err = ctx.GetTable(node[i].TableName)
		if err != nil {
			return t.nodeFromTables(qTables), err
		}

		qTables[i] = anode.Table{
			Def:      tableDef,
			Alias:    node[i].TableAlias,
			QTableId: ctx.CurrentQuery.NextQTableId(tableDef.TableId()),
		}
	}

	return t.nodeFromTables(qTables), nil
}

func (t tables) nodeFromTables(tables []anode.Table) anode.Tables {
	tablesNode := anode.Tables{Tables: tables}
	tablesNode.Init()

	return tablesNode
}
