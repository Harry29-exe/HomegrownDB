package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hg"
	table "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	txCtx tx.Tx,
	store hg.DBStore,
) ExCtx {
	cache, rteMap := createCache(stmt.Tables, store.TableStore())
	return &executionCtx{
		Stmt:     stmt,
		Buff:     store.SharedBuffer(),
		FsmStore: store.FsmStore(),
		Tables:   cache,
		Tx:       txCtx,
		rteMap:   rteMap,
	}
}

func createCache(rteList []node.RangeTableEntry, store table.Store) (table.Cache, map[node.RteID]node.RangeTableEntry) {
	cache := map[table.Id]table.RDefinition{}
	rteMap := map[node.RteID]node.RangeTableEntry{}
	for _, rte := range rteList {
		if rte.Kind == node.RteRelation {
			tab := store.AccessTable(rte.TableId, rte.LockMode)
			cache[rte.TableId] = tab
			rte.Ref = tab
		}
		rteMap[rte.Id] = rte
	}
	return cache, rteMap
}

type executionCtx struct {
	Stmt     node.PlanedStmt
	Buff     buffer.SharedBuffer
	FsmStore fsm.Store
	Tables   table.Cache
	Tx       tx.Tx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}
