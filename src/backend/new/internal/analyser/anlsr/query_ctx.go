package anlsr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser/tokenizer"
	"HomegrownDB/backend/new/internal/sqlerr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
)

type Ctx = *queryCtx

func NewQCtx(store table.Store) Ctx {
	return &queryCtx{
		RteIdCounter: appsync.NewSyncCounter[node.RteID](0),
		TableStore:   store,
		TableCache:   map[relation.ID]table.Definition{},
		TableIdMap:   map[string]relation.ID{},
	}
}

type queryCtx struct {
	RteIdCounter appsync.SyncCounter[node.RteID]

	TableStore table.Store
	TableCache map[relation.ID]table.Definition
	TableIdMap map[string]relation.ID // TableIdMap map[tableName] = tableId

	TokenSrc tokenizer.TokenSource
}

func (c Ctx) GetTableById(id relation.ID) table.RDefinition {
	cachedTable, ok := c.TableCache[id]
	if ok {
		return cachedTable
	}

	tab := c.TableStore.AccessTable(id, table.RLockMode)
	c.TableCache[id] = tab
	return tab
}

func (c Ctx) GetTable(name string) (table.RDefinition, error) {
	tableId, ok := c.TableIdMap[name]
	if ok {
		return c.TableCache[tableId], nil
	}

	tableId = c.TableStore.FindTable(name)
	if tableId == relation.InvalidRelId {
		return nil, sqlerr.NewNoTableWithName(name)
	}

	tableDef := c.TableStore.AccessTable(tableId, table.RLockMode)
	c.TableCache[tableId] = tableDef
	return tableDef, nil
}
