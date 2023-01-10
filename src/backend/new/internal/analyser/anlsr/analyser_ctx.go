package anlsr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/sqlerr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/relation"
	table2 "HomegrownDB/dbsystem/relation/table"
)

type Ctx = *ctx

func NewCtx(store table2.Store) Ctx {
	return &ctx{
		RteIdCounter: appsync.NewSimpleCounter[node.RteID](0),

		TableStore: store,
		TableCache: map[relation.ID]table2.Definition{},
		TableIdMap: map[string]relation.ID{},
	}
}

type ctx struct {
	RteIdCounter RteIdCounter

	TableStore table2.Store
	TableCache map[relation.ID]table2.Definition
	TableIdMap map[string]relation.ID // TableIdMap map[tableName] = tableId
}

func (c Ctx) GetTableById(id relation.ID) table2.RDefinition {
	cachedTable, ok := c.TableCache[id]
	if ok {
		return cachedTable
	}

	tab := c.TableStore.AccessTable(id, table2.RLockMode)
	c.TableCache[id] = tab
	return tab
}

func (c Ctx) GetTable(name string) (table2.RDefinition, error) {
	tableId, ok := c.TableIdMap[name]
	if ok {
		return c.TableCache[tableId], nil
	}

	tableId = c.TableStore.FindTable(name)
	if tableId == relation.InvalidRelId {
		return nil, sqlerr.NewNoTableWithNameErr(name)
	}

	tableDef := c.TableStore.AccessTable(tableId, table2.RLockMode)
	c.TableCache[tableId] = tableDef
	return tableDef, nil
}

type RteIdCounter = appsync.SimpleSyncCounter[node.RteID]