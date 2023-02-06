package anlsr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/reldef"
	table2 "HomegrownDB/dbsystem/reldef/tabdef"
)

type Ctx = *ctx

func NewCtx(store table.Store) Ctx {
	return &ctx{
		RteIdCounter: appsync.NewSimpleCounter[node.RteID](0),

		TableStore: store,
		TableCache: map[reldef.OID]table2.Definition{},
		TableIdMap: map[string]reldef.OID{},
	}
}

type ctx struct {
	RteIdCounter RteIdCounter

	TableStore table.Store
	TableCache map[reldef.OID]table2.Definition
	TableIdMap map[string]reldef.OID // TableIdMap map[tableName] = tableId
}

func (c Ctx) GetTableById(id reldef.OID) table2.RDefinition {
	cachedTable, ok := c.TableCache[id]
	if ok {
		return cachedTable
	}

	tab := c.TableStore.AccessTable(id, table.RLockMode)
	c.TableCache[id] = tab
	return tab
}

func (c Ctx) GetTable(name string) (table2.RDefinition, error) {
	tableId, ok := c.TableIdMap[name]
	if ok {
		return c.TableCache[tableId], nil
	}

	tableId = c.TableStore.FindTable(name)
	if tableId == reldef.InvalidRelId {
		return nil, sqlerr.NewNoTableWithNameErr(name)
	}

	tableDef := c.TableStore.AccessTable(tableId, table.RLockMode)
	c.TableCache[tableId] = tableDef
	return tableDef, nil
}

type RteIdCounter = appsync.SimpleSyncCounter[node.RteID]
