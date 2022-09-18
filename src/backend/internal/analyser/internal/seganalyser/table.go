package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/table"
)

var Tables = tables{}

type tables struct{}

func (t tables) Analise(node pnode.TablesNode, ctx *internal.AnalyserCtx) (anode.Tables, error) {
	tablesCount := len(node.Tables)
	qTables := make([]anode.Table, tablesCount)

	var tableDef table.Definition
	var err error
	for i := 0; i < tablesCount; i++ {
		tableDef, err = ctx.GetTable(node.Tables[i].TableName)
		if err != nil {
			return t.nodeFromTables(qTables), err
		}

		qTables[i] = anode.Table{
			Def:      tableDef,
			Alias:    node.Tables[i].TableAlias,
			QTableId: ctx.NextQTableId(tableDef.TableId()),
		}
	}

	return t.nodeFromTables(qTables), nil
}

func (t tables) nodeFromTables(tables []anode.Table) anode.Tables {
	tablesNode := anode.Tables{Tables: tables}
	tablesNode.Init()

	return tablesNode
}