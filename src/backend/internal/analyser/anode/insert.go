package anode

import (
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/dbsystem/schema/column"
)

type Insert struct {
	Table   qctx.QTableId
	Columns []column.Id

	// Rows if insert has 'VALUE' or 'VALUES' keyword,
	// e.g. INSERT INTO uses VALUES (...
	// then this field contains all information to fill rows
	Rows []InsertRow
	// Expression
	Expression any
}
