package exinfr

import (
	node "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/hg"
	relation2 "HomegrownDB/dbsystem/reldef"
	table2 "HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type ExCtx = *executionCtx

func NewExCtx(
	stmt node.PlanedStmt,
	txCtx tx.Tx,
	container hg.ExecutionContainer,
) ExCtx {
	cache, rteMap := createCache(stmt.Tables, container.RelationManager)
	return &executionCtx{
		Stmt:       stmt,
		Buff:       container.SharedBuffer,
		FsmStore:   container.FsmStore,
		Tables:     cache,
		TableStore: container.RelationManager,
		Tx:         txCtx,
		rteMap:     rteMap,
	}
}

func createCache(rteList []node.RangeTableEntry, store relation.Manager) (table2.Cache, map[node.RteID]node.RangeTableEntry) {
	cache := map[table2.Id]table2.RDefinition{}
	rteMap := map[node.RteID]node.RangeTableEntry{}
	for _, rte := range rteList {
		if rte.Kind == node.RteRelation {
			store.Lock(rte.TableId, rte.LockMode)
			rel := store.GetByOID(rte.TableId)
			if rel.Kind() != relation2.TypeTable {
				panic("illegal type")
			}
			tableDef := rel.(table2.RDefinition)
			cache[rte.TableId] = tableDef
			rte.Ref = tableDef
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
	TableStore relation.Manager
	Tx         tx.Tx

	rteMap map[node.RteID]node.RangeTableEntry
}

func (e ExCtx) GetRTE(id node.RteID) node.RangeTableEntry {
	return e.rteMap[id]
}

func (e ExCtx) Close() error {
	//for i, tableDef := range e.Tables {
	//	e.RelationManager.AccessTable()
	//}
	//todo implement me
	panic("Not implemented")
}
