package tx

import (
	"HomegrownDB/dbsystem/schema/table"
	"strings"
)

type QueryCtx struct {
	QueryTokens []string

	nextQTableId QTableId
	// qTableIdTableIdMap slice functions here as map as x[anode.QTableId] = table.Id
	qTableIdTableIdMap []table.Id
}

type QTableId = uint16

func (c *QueryCtx) Reconstruct(startToken, endToken uint32) string {
	strBuilder := strings.Builder{}
	for i := startToken; i < endToken; i++ {
		strBuilder.WriteString(c.QueryTokens[i])
	}
	return strBuilder.String()
}

func (c *QueryCtx) GetTableId(qTableId QTableId) table.Id {
	return c.qTableIdTableIdMap[qTableId]
}

func (c *QueryCtx) NextQTableId(tableId table.Id) (id QTableId) {
	id = c.nextQTableId
	c.qTableIdTableIdMap = append(c.qTableIdTableIdMap, tableId)
	c.nextQTableId++
	if int(id) != len(c.qTableIdTableIdMap) {
		panic("illegal state")
	}

	return id
}
