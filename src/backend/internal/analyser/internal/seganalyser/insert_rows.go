package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
)

var InsertRows = insertRows{}

type insertRows struct{}

func (ir insertRows) Analyse(
	rows []pnode.InsertingRow,
	columnIds []qctx.QColumnId,
	ctx qctx.QueryCtx,
) ([]anode.InsertRow, error) {

	rowsNode := make([]anode.InsertRow, len(rows))

	tableDef := ctx.QTCtx.GetTableByQTableId(columnIds[0].QTableId)
	columns := make([]column.Def, len(columnIds))
	for i := 0; i < len(columnIds); i++ {
		columns[i] = tableDef.Column(columnIds[i].ColOrder)
	}

	for i, row := range rows {
		analysedRow, err := ir.analyseRow(row, columns)
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
