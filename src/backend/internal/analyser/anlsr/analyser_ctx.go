package anlsr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/relation/table"
)

type Ctx = *ctx

func NewCtx(store table.Store) Ctx {
	return &ctx{
		RteIdCounter: appsync.NewSimpleCounter[node.RteID](0),

		TableStore: store,
		TableCache: map[relation.OID]table.Definition{},
		TableIdMap: map[string]relation.OID{},
	}
}

type ctx struct {
	RteIdCounter RteIdCounter

	TableStore table.Store
	TableCache map[relation.OID]table.Definition
	TableIdMap map[string]relation.OID // TableIdMap map[tableName] = tableId
}

func (c Ctx) GetTableById(id relation.OID) table.RDefinition {
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
		return nil, sqlerr.NewNoTableWithNameErr(name)
	}

	tableDef := c.TableStore.AccessTable(tableId, table.RLockMode)
	c.TableCache[tableId] = tableDef
	return tableDef, nil
}

type RteIdCounter = appsync.SimpleSyncCounter[node.RteID]
