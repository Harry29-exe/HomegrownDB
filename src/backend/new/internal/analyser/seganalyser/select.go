package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var Select = _select{}

type _select struct{}

func (s _select) Analyse(stmt pnode.SelectStmt, ctx query.Ctx) node.Query {
	query := node.NewQuery()
	query.Command = node.CommandTypeSelect

}
