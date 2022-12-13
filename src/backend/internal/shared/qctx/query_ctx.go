package qctx

import (
	"HomegrownDB/dbsystem/schema/table"
	"strings"
)

func NewQueryCtx(tableStore table.Store) QueryCtx {
	return &queryCtx{
		QueryTokens: nil,
		QTCtx:       NewQTableCtx(tableStore),
	}
}

type QueryCtx = *queryCtx

type queryCtx struct {
	QueryTokens []string
	// QTableCtx is proxy between compile phase and table.Store,
	// it will any locks table accessed during query compilation
	QTCtx QTableCtx
	//QField QFieldCtx

	fieldAliases map[QColumnId]string
}

func (c QueryCtx) Reconstruct(startToken, endToken uint32) string {
	strBuilder := strings.Builder{}
	for i := startToken; i < endToken; i++ {
		strBuilder.WriteString(c.QueryTokens[i])
	}
	return strBuilder.String()
}

func (c QueryCtx) AddFieldAlias(alias string, fieldId QColumnId) {
	c.fieldAliases[fieldId] = alias
}

func (c QueryCtx) GetAlias(fieldId QColumnId) (alias string, ok bool) {
	alias, ok = c.fieldAliases[fieldId]
	return
}
