package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/tx"
)

var InsertRows = insertRows{}

type insertRows struct{}

func (ir insertRows) Analyse(
	rows []pnode.InsertingRow,
	cols []column.Def,
	ctx *tx.Ctx,
) ([]anode.InsertRow, error) {

	rowsNode := make([]anode.InsertRow, len(rows))

	for i, row := range rows {
		analysedRow, err := ir.analyseRow(row, cols)
		if err != nil {
			return nil, err
		}

		rowsNode[i] = analysedRow
	}

	return rowsNode, nil
}

func (ir insertRows) analyseRow(row pnode.InsertingRow, columns []column.Def) (anode.InsertRow, error) {
	aRow := anode.InsertRow{Fields: make([]anode.InsertField, 0, len(columns))}

	for i, field := range row.Fields {
		if val := field.Value; val != nil {
			cval, err := ctype.ConvInput(val.V, columns[i].Type())
			if err != nil {
				return anode.InsertRow{}, err
			}

			aRow.Fields = append(aRow.Fields, anode.InsertField{Value: cval})
		} else {
			panic("expression support not yet implemented in insert")
		}
	}

	return aRow, nil
}
