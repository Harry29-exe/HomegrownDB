package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	buff buffer.SharedBuffer,
	store table.Store,
	txCtx *tx.Ctx,
) ExCtx {
	cache := createCache(stmt.Tables, store)
	return &executionCtx{
		Stmt:   stmt,
		Buff:   buff,
		Tables: cache,
		TxCtx:  txCtx,
	}
}

func createCache(rteList []node.RangeTableEntry, store table.Store) map[table.Id]table.RDefinition {
	cache := map[table.Id]table.RDefinition{}
	for _, rte := range rteList {
		if rte.Kind == node.RteRelation {
			tab := store.AccessTable(rte.TableId, rte.LockMode)
			cache[rte.TableId] = tab
		}
	}
	return cache
}

type executionCtx struct {
	Stmt   node.PlanedStmt
	Buff   buffer.SharedBuffer
	Tables table.Cache
	TxCtx  *tx.Ctx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}
