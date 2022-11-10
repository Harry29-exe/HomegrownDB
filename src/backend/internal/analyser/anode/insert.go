package anode

import "HomegrownDB/dbsystem/schema/column"

type Insert struct {
	Table   Table
	Columns []column.Def

	// Rows if insert has 'VALUE' or 'VALUES' keyword,
	// e.g. INSERT INTO uses VALUES (...
	// then this field contains all information to fill rows
	Rows []InsertRow
	// Expression
	Expression any
}
