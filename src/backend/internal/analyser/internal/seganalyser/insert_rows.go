package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/tx"
)

var InsertRows = insertRows{}

type insertRows struct{}

func (receiver insertRows) Analyse(
	rows []pnode.InsertingRow,
	pattern ColumnTypesPattern,
	ctx *tx.Ctx,
) (*anode.InsertRows, error) {

	insertRows := anode.NewInsertRows(uint(len(rows)), uint16(len(pattern.Types)))

	for _, row := range rows {
		if err := pattern.RowMatches(row, ctx); err != nil {
			return nil, err
		}

		for i, val := range row.Values {
			insertRows.PutValue(val.ConvertTo(pattern.Types[i]))
		}
	}

	return insertRows, nil
}
