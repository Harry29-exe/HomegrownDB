package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hg"
	table2 "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	txCtx *tx.Ctx,
	store hg.DBStore,
) ExCtx {
	cache := createCache(stmt.Tables, store.TableStore())
	return &executionCtx{
		Stmt:     stmt,
		Buff:     store.SharedBuffer(),
		FsmStore: store.FsmStore(),
		Tables:   cache,
		TxCtx:    txCtx,
	}
}

func createCache(rteList []node.RangeTableEntry, store table2.Store) map[table2.Id]table2.RDefinition {
	cache := map[table2.Id]table2.RDefinition{}
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
	Tables   table2.Cache
	TxCtx    *tx.Ctx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}
