package internal

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/stores"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(node pnode.SelectNode, store stores.Tables, ctx analyser.Ctx) (anode.Select, error) {
	selectNode := anode.Select{}
	tablesNode, err := Tables.Analise(node.Tables, store, ctx)
	if err != nil {
		return selectNode, err
	}
	selectNode.Tables = tablesNode

	fieldsNode, err := Fields.Analyse(node.Fields, tablesNode)
	if err != nil {
		return selectNode, err
	}
	selectNode.Fields = fieldsNode

	return selectNode, err
}
