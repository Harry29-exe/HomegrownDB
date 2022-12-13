package tx

import (
	"HomegrownDB/dbsystem/schema/table"
)

type Ctx struct {
	rLockedTables map[table.Id]bool
	tableStore    table.RDefsStore

	Info *InfoCtx
}

func NewContext(txId Id, tableStore table.RDefsStore) *Ctx {
	return &Ctx{
		rLockedTables: make(map[table.Id]bool, 20),
		tableStore:    tableStore,
		Info:          NewInfoCtx(txId),
	}
}
