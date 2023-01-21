package exinfr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hg"
	table "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node2.PlanedStmt,
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

func createCache(rteList []node2.RangeTableEntry, store table.Store) (table.Cache, map[node2.RteID]node2.RangeTableEntry) {
	cache := map[table.Id]table.RDefinition{}
	rteMap := map[node2.RteID]node2.RangeTableEntry{}
	for _, rte := range rteList {
		if rte.Kind == node2.RteRelation {
			tab := store.AccessTable(rte.TableId, rte.LockMode)
			cache[rte.TableId] = tab
			rte.Ref = tab
		}
		rteMap[rte.Id] = rte
	}
	return cache, rteMap
}

type executionCtx struct {
	Stmt     node2.PlanedStmt
	Buff     buffer.SharedBuffer
	FsmStore fsm.Store
	Tables   table.Cache
	Tx       tx.Tx

	rteMap map[node2.RteID]node2.RangeTableEntry
}

func (e ExCtx) GetRTE(id node2.RteID) node2.RangeTableEntry {
	return e.rteMap[id]
}
