package qctx

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"errors"
)

func NewQTableCtx(store table.Store) QTableCtx {
	return &qTableCtx{
		lockedTables: map[table.Id]table.Definition{},
		qTableIdMap:  make([]table.Id, 0, 10),
		aliasMap:     map[string]QTableId{},
		tableStore:   store,
	}
}

// QTableCtx is proxy between analyser/planner and table store,
// it rlocks tables, so they can't be modified during analysing query.
// It is not designed to support concurrent access.
type QTableCtx = *qTableCtx

type qTableCtx struct {
	lockedTables map[table.Id]table.Definition
	// qTableIdMap is map: qTableIdMap[QTableId] = table.Id
	qTableIdMap []table.Id
	aliasMap    map[string]QTableId

	tableStore table.Store
}

// GetTable returns table with given id, if table was
// accessed before it is returned from cache,
// otherwise it is rlocked, cached and returned
func (q QTableCtx) GetTable(id table.Id) table.RDefinition {
	tableDef, ok := q.lockedTables[id]
	if ok {
		return tableDef
	}
	tableDef = q.tableStore.AccessTable(id, table.RLockMode)
	q.lockedTables[id] = tableDef
	return tableDef
}

// GetTableByName returns table with given name, if table was
// accessed before it is returned from cache, otherwise
// if it exists, it is rlocked, cached and returned, otherwise error is returned
func (q QTableCtx) GetTableByName(name string) (table.RDefinition, error) {
	tableId := q.tableStore.FindTable(name)
	if tableId == relation.InvalidRelId {
		return nil, errors.New("no table with such name") //todo better error
	}

	return q.GetTable(tableId), nil
}

func (q QTableCtx) GetTableByQTableId(qTableId QTableId) table.RDefinition {
	return q.lockedTables[q.qTableIdMap[qTableId]]
}

func (q QTableCtx) GetOrCreateQTableId(alias string, def table.RDefinition) (QTableId, error) {
	//todo implement me
	panic("Not implemented")
	//qid, ok := q.aliasMap[alias]
	//if !ok {
	//	qid = QTableId(len(q.qTableIdMap))
	//	q.aliasMap[alias] = qid
	//	q.qTableIdMap = append(q.qTableIdMap, qid)
	//} else {
	//	tableId := q.qTableIdMap[qid]
	//	if tableId != def.TableId() {
	//		return InvalidQTableId, errors.New("") //todo better error
	//	}
	//}
	//
	//return qid, nil
}

func (q QTableCtx) GetQTableId(alias string) QTableId {
	qid, ok := q.aliasMap[alias]
	if !ok {
		return InvalidQTableId
	}
	return qid
}
