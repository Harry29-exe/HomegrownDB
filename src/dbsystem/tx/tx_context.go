package tx

import (
	table2 "HomegrownDB/dbsystem/relation/table"
)

type Ctx struct {
	rLockedTables map[table2.Id]bool
	tableStore    table2.RDefsStore

	Info *InfoCtx
}

func NewContext(txId Id, tableStore table2.RDefsStore) *Ctx {
	return &Ctx{
		rLockedTables: make(map[table2.Id]bool, 20),
		tableStore:    tableStore,
		Info:          NewInfoCtx(txId),
	}
}
