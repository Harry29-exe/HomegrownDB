package exinfr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/hg"
	"HomegrownDB/dbsystem/reldef"
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

type tableCache = map[reldef.OID]reldef.TableRDefinition

func createCache(rteList []node.RangeTableEntry, store relation.Manager) (tableCache, map[node.RteID]node.RangeTableEntry) {
	cache := map[reldef.OID]reldef.TableRDefinition{}
	rteMap := map[node.RteID]node.RangeTableEntry{}
	for _, rte := range rteList {
		if rte.Kind == node.RteRelation {
			rel := store.Access(rte.TableId, rte.LockMode)
			if rel.Kind() != reldef.TypeTable {
				panic("illegal type")
			}
			tableDef := rel.(reldef.TableRDefinition)
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
	Tables     tableCache
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
