package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Analyse(stmt pnode.InsertStmt, ctx anlsr.Ctx) (node.Query, error) {
	//todo implement me
	panic("Not implemented")
}
