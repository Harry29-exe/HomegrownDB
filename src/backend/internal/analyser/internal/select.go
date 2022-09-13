package internal

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(node pnode.SelectNode, ctx *AnalyserCtx) (anode.Select, error) {
	selectNode := anode.Select{}
	tablesNode, err := Tables.Analise(node.Tables, ctx)
	if err != nil {
		return selectNode, err
	}
	selectNode.Tables = tablesNode

	fieldsNode, err := Fields.Analyse(node.Fields, tablesNode, ctx)
	if err != nil {
		return selectNode, err
	}
	selectNode.Fields = fieldsNode

	return selectNode, err
}
