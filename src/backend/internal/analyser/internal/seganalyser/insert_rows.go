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

func (receiver insertRows) Analyse(
	rows []pnode.InsertingRow,
	cols []column.Def,
	ctx *tx.Ctx,
) (*anode.InsertRows, error) {

	rowsNode := anode.NewInsertRows(uint(len(rows)), uint16(len(cols)))

	for _, row := range rows {
		for i, val := range row.Values {
			cval, err := ctype.ConvInput(val, cols[i].Type())
			if err != nil {
				return nil, err
			}

			rowsNode.PutValue(cval)
		}
	}

	return rowsNode, nil
}
