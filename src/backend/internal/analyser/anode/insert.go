package anode

import (
	"HomegrownDB/backend/internal/shared/qctx"
)

type Insert struct {
	Table   qctx.QTableId
	Columns []qctx.QColumnId

	// Rows if insert has 'VALUE' or 'VALUES' keyword,
	// e.g. INSERT INTO uses VALUES (...
	// then this field contains all information to fill rows
	Rows []InsertRow
	// Expression
	Expression any
}
