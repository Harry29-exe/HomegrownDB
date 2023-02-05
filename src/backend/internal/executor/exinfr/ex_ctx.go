package exinfr

import (
	node "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	table2 "HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/hg/di"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	txCtx tx.Tx,
	container di.ExecutionContainer,
) ExCtx {
	cache, rteMap := createCache(stmt.Tables, container.TableStore)
	return &executionCtx{
		Stmt:       stmt,
		Buff:       container.SharedBuffer,
		FsmStore:   container.FsmStore,
		Tables:     cache,
		TableStore: container.TableStore,
		Tx:         txCtx,
		rteMap:     rteMap,
	}
}

func createCache(rteList []node.RangeTableEntry, store table2.Store) (table2.Cache, map[node.RteID]node.RangeTableEntry) {
	cache := map[table2.Id]table2.RDefinition{}
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
	Stmt       node.PlanedStmt
	Buff       buffer.SharedBuffer
	FsmStore   fsm.Store
	Tables     table2.Cache
	TableStore table2.Store
	Tx         tx.Tx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}

func (e ExCtx) Close() error {
	//for i, tableDef := range e.Tables {
	//	e.TableStore.AccessTable()
	//}
	//todo implement me
	panic("Not implemented")
}
