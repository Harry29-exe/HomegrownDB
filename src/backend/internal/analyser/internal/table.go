package internal

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
)

var Tables = tables{}

type tables struct{}

func (t tables) Analise(node pnode.TablesNode, store stores.Tables, ctx analyser.Ctx) (anode.Tables, error) {
	tablesCount := len(node.Tables)
	qTables := make([]anode.Table, tablesCount)

	var tableDef table.Definition
	var err error
	for i := 0; i < tablesCount; i++ {
		tableDef, err = store.GetTable(node.Tables[i].TableName)
		if err != nil {
			return t.nodeFromTables(qTables), err
		}

		qTables[i] = anode.Table{
			Table: tableDef,
			Alias: node.Tables[i].TableAlias,
		}
	}

	return t.nodeFromTables(qTables), nil
}

func (t tables) nodeFromTables(tables []anode.Table) anode.Tables {
	tablesNode := anode.Tables{Tables: tables}
	tablesNode.Init()

	return tablesNode
}
