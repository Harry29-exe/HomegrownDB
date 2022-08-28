package analyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
)

func AnaliseTables(node pnode.TablesNode, store stores.Tables) (anode.Tables, error) {
	tablesCount := len(node.Tables)
	newNode := anode.Tables{Tables: make([]anode.Table, tablesCount)}

	var tableDef table.Definition
	var err error
	for i := 0; i < tablesCount; i++ {
		tableDef, err = store.GetTable(node.Tables[i].TableName)
		if err != nil {
			return newNode, err
		}

		newNode.Tables[i] = anode.Table{
			Table: tableDef,
			Alias: node.Tables[i].TableAlias,
		}
	}

	newNode.Init()
	return newNode, nil
}
