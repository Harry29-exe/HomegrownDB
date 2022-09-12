package analyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/dbsystem/schema/table"
)

type tablesCtx struct {
	nextQtableId anode.QtableId
	// qtableIdTableIdMap slice functions here as map as x[anode.QtableId] = table.Id
	qtableIdTableIdMap []table.Id
}

func (t *tablesCtx) NextQtableId(tableId table.Id) (id anode.QtableId) {
	id = t.nextQtableId
	t.qtableIdTableIdMap = append(t.qtableIdTableIdMap, tableId)
	t.nextQtableId++
	if int(id) != len(t.qtableIdTableIdMap) {
		panic("illegal state")
	}

	return id
}

func (t *tablesCtx) GetTableId(qtableId anode.QtableId) table.Id {
	return t.qtableIdTableIdMap[qtableId]
}
