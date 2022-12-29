package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	txCtx *tx.Ctx,
	store dbsystem.DBSystem,
) ExCtx {
	cache := createCache(stmt.Tables, store.TableStore())
	return &executionCtx{
		Stmt:     stmt,
		Buff:     store.Buffer(),
		FsmStore: store.FsmStore(),
		Tables:   cache,
		TxCtx:    txCtx,
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
	Stmt     node.PlanedStmt
	Buff     buffer.SharedBuffer
	FsmStore fsm.Store
	Tables   table.Cache
	TxCtx    *tx.Ctx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}
