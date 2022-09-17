package internal

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/dbsystem/schema/table"
)

type tablesCtx struct {
	nextQTableId anode.QTableId
	// qTableIdTableIdMap slice functions here as map as x[anode.QTableId] = table.Id
	qTableIdTableIdMap []table.Id
}

func (t *tablesCtx) NextQTableId(tableId table.Id) (id anode.QTableId) {
	id = t.nextQTableId
	t.qTableIdTableIdMap = append(t.qTableIdTableIdMap, tableId)
	t.nextQTableId++
	if int(id) != len(t.qTableIdTableIdMap) {
		panic("illegal state")
	}

	return id
}

func (t *tablesCtx) GetTableId(qTableId anode.QTableId) table.Id {
	return t.qTableIdTableIdMap[qTableId]
}
