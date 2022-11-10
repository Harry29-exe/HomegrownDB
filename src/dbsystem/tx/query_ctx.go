package tx

import (
	"HomegrownDB/dbsystem/schema/table"
	"strings"
)

func NewQueryCtx() QueryCtx {
	return &queryCtx{
		QueryTokens:        nil,
		nextQTableId:       0,
		qTableIdTableIdMap: make([]table.Id, 0, 20),
	}
}

type QueryCtx = *queryCtx

type queryCtx struct {
	QueryTokens []string

	nextQTableId QTableId
	// qTableIdTableIdMap slice functions here as map as x[anode.QTableId] = table.Id
	qTableIdTableIdMap []table.Id
	qTableAliasMap     map[string]QTableId
}

type QTableId = uint16

func (c QueryCtx) Reconstruct(startToken, endToken uint32) string {
	strBuilder := strings.Builder{}
	for i := startToken; i < endToken; i++ {
		strBuilder.WriteString(c.QueryTokens[i])
	}
	return strBuilder.String()
}

func (c QueryCtx) GetTableId(qTableId QTableId) table.Id {
	return c.qTableIdTableIdMap[qTableId]
}

func (c QueryCtx) GetOrCreateQTableId(alias string, tableId table.Id) QTableId {
	qtableId, ok := c.qTableAliasMap[alias]
	if ok {
		return qtableId
	}

	qTableId := QTableId(len(c.qTableIdTableIdMap))
	c.qTableIdTableIdMap = append(c.qTableIdTableIdMap, tableId)
	c.qTableAliasMap[alias] = qTableId
	return qtableId
}

func (c QueryCtx) GetQTableId(alias string) (id QTableId, ok bool) {
	id, ok = c.qTableAliasMap[alias]
	return
}

//func (c QueryCtx) NextQTableId(tableId table.Id) (id QTableId) {
//	id = c.nextQTableId
//	c.qTableIdTableIdMap = append(c.qTableIdTableIdMap, tableId)
//	c.nextQTableId++
//	if int(c.nextQTableId) != len(c.qTableIdTableIdMap) {
//		panic("illegal state")
//	}
//
//	return id
//}
