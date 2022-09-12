package analyser

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/stores"
	"HomegrownDB/dbsystem/tx"
)

func NewAnalyserCtx(txCtx tx.Ctx, tables stores.RTablesDefs) *Ctx {
	return &Ctx{
		tableStore: tables,
		txCtx:      txCtx,
		tablesCtx: tablesCtx{
			nextQtableId:       0,
			qtableIdTableIdMap: make([]table.Id, 0, 10),
		},
	}
}

type Ctx struct {
	tableStore stores.RTablesDefs
	txCtx      tx.Ctx
	tablesCtx
}

func (c *Ctx) GetTable(name string) (table.Definition, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Ctx) Table(id table.Id) table.Definition {
	//TODO implement me
	panic("implement me")
}

var (
	_ stores.RTablesDefs = (*Ctx)(nil)
)
