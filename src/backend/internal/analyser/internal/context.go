package internal

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

func NewAnalyserCtx(txCtx tx.Ctx, tables table.RDefsStore) *AnalyserCtx {
	return &AnalyserCtx{
		tableStore: tables,
		txCtx:      txCtx,
		tablesCtx: tablesCtx{
			nextQTableId:       0,
			qTableIdTableIdMap: make([]table.Id, 0, 10),
		},
	}
}

var (
	_ table.RDefsStore = (*AnalyserCtx)(nil)
)

type AnalyserCtx struct {
	tableStore table.RDefsStore
	txCtx      tx.Ctx
	tablesCtx
}

func (c *AnalyserCtx) GetTable(name string) (table.Definition, error) {
	//TODO implement me
	panic("implement me")
}

func (c *AnalyserCtx) Table(id table.Id) table.Definition {
	//TODO implement me
	panic("implement me")
}
