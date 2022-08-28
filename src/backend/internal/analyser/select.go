package analyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/stores"
)

func AnalyseSelect(node pnode.SelectNode, store stores.Tables) (anode.Select, error) {
	selectNode := anode.Select{}
	tables, err := AnaliseTables(node.Tables, store)
	if err != nil {
		return selectNode, err
	}
	selectNode.Tables = tables

	fields, err := AnalyseFields(node.Fields, tables)
	if err != nil {
		return selectNode, err
	}
	selectNode.Fields = fields

	return selectNode, err
}
