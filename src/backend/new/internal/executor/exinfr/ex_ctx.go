package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/table"
)

type ExCtx = *executionCtx

func NewExCtx(buff buffer.SharedBuffer, store table.Store, rteList []node.RangeTableEntry) ExCtx {
	cache := createCache(rteList, store)
	return &executionCtx{
		Buff:   buff,
		Tables: cache,
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
	Buff   buffer.SharedBuffer
	Tables table.Cache
}
