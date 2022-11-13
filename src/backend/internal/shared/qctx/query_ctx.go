package qctx

import (
	"HomegrownDB/dbsystem/schema/table"
	"strings"
)

func NewQueryCtx() QueryCtx {
	return &queryCtx{
		QueryTokens: nil,

		qTableIdTableIdMap: make([]table.Id, 0, 20),
		qTableAliasMap:     map[string]QTableId{},
	}
}

type QueryCtx = *queryCtx

type queryCtx struct {
	QueryTokens []string
	// QTableCtx is proxy between compile phase and table.Store,
	// it will any locks table accessed during query compilation
	QTCtx QTableCtx
	//QField QFieldCtx

	// qTableIdTableIdMap slice functions here as map as x[anode.QTableId] = table.Id
	qTableIdTableIdMap []table.Id
	qTableAliasMap     map[string]QTableId

	fieldAliases map[QFieldId]string
}

func (c QueryCtx) Reconstruct(startToken, endToken uint32) string {
	strBuilder := strings.Builder{}
	for i := startToken; i < endToken; i++ {
		strBuilder.WriteString(c.QueryTokens[i])
	}
	return strBuilder.String()
}

func (c QueryCtx) AddFieldAlias(alias string, fieldId QFieldId) {
	c.fieldAliases[fieldId] = alias
}

func (c QueryCtx) GetAlias(fieldId QFieldId) (alias string, ok bool) {
	alias, ok = c.fieldAliases[fieldId]
	return
}
