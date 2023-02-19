package anlctx

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/reldef"
	table "HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/lib/datastructs/appsync"
)

type Ctx = *ctx

func NewCtx(store relation.AccessMngr) Ctx {
	return &ctx{
		RteIdCounter: appsync.NewSimpleCounter[node.RteID](0),

		TableStore: store,
		TableCache: map[reldef.OID]table.TableRDefinition{},
		TableIdMap: map[string]reldef.OID{},
	}
}

type ctx struct {
	RteIdCounter RteIdCounter

	TableStore relation.AccessMngr
	TableCache map[reldef.OID]table.TableRDefinition
	TableIdMap map[string]reldef.OID // TableIdMap map[tableName] = tableId
}

func (c Ctx) GetTableById(id reldef.OID) table.TableRDefinition {
	cachedTable, ok := c.TableCache[id]
	if ok {
		return cachedTable
	}

	rel := c.TableStore.Access(id, relation.LockRead)
	if rel.Kind() != reldef.TypeTable {

	}
	tab := rel.(table.TableRDefinition)
	c.TableCache[id] = tab
	return tab
}

func (c Ctx) GetTable(name string) (table.TableRDefinition, error) {
	tableId, ok := c.TableIdMap[name]
	if ok {
		return c.TableCache[tableId], nil
	}

	tableId = c.TableStore.FindByName(name)
	if tableId == reldef.InvalidRelId {
		return nil, sqlerr.NewNoTableWithNameErr(name)
	}

	rel := c.TableStore.Access(tableId, relation.LockRead)
	if rel.Kind() != reldef.TypeTable {
		return nil, sqlerr.NewNoTableWithNameErr(name)
	}
	tableDef := rel.(table.TableDefinition)
	c.TableCache[tableId] = tableDef
	return tableDef, nil
}

type RteIdCounter = appsync.SimpleSyncCounter[node.RteID]
