package qctx

import (
	"HomegrownDB/dbsystem/schema/table"
	"errors"
)

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

func (q QTableCtx) GetTable(id table.Id) table.RDefinition {
	tableDef, ok := q.lockedTables[id]
	if ok {
		return tableDef
	}
	tableDef = q.tableStore.AccessTable(id, table.RLockMode)
	q.lockedTables[id] = tableDef
	return tableDef
}

func (q QTableCtx) GetTableByName(name string) (table.RDefinition, error) {
	tableId := q.tableStore.FindTable(name)
	if tableId == table.InvalidTableId {
		return nil, errors.New("no table with such name") //todo better error
	}

	return q.GetTable(tableId), nil
}

func (q QTableCtx) GetTableByQTableId(qTableId QTableId) table.Definition {
	return q.lockedTables[q.qTableIdMap[qTableId]]
}

func (q QTableCtx) GetOrCreateQTableId(alias string, def table.RDefinition) (QTableId, error) {
	qid, ok := q.aliasMap[alias]
	if !ok {
		qid = QTableId(len(q.qTableIdMap))
		q.aliasMap[alias] = qid
		q.qTableIdMap = append(q.qTableIdMap, qid)
	} else {
		tableId := q.qTableIdMap[qid]
		if tableId != def.TableId() {
			return InvalidQTableId, errors.New("") //todo better error
		}
	}

	return qid, nil
}

func (q QTableCtx) GetQTableId(alias string) QTableId {
	qid, ok := q.aliasMap[alias]
	if !ok {
		return InvalidQTableId
	}
	return qid
}
