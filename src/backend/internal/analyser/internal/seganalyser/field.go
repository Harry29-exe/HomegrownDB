package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"errors"
)

var Fields = fields{}

type fields struct{}

func (f fields) Analyse(fieldNodes []pnode.FieldNode, tables []qctx.QTableId, ctx qctx.QueryCtx) (anode.SelectFields, error) {
	fieldsCount := len(fieldNodes)
	fieldsNode := make([]anode.SelectField, fieldsCount)

	for i, field := range fieldNodes {
		qTableId := ctx.QTCtx.GetQTableId(field.TableAlias)
		if qTableId == qctx.InvalidQTableId {
			return anode.SelectFields{}, errors.New("") // todo better message
		}

		table := ctx.QTCtx.GetTableByQTableId(qTableId)
		colOrder, ok := table.ColumnOrder(field.FieldName)
		if !ok {
			return anode.SelectFields{}, errors.New("") // todo better message
		}

		fieldsNode[i] = anode.SelectField{
			Table:  qTableId,
			Column: colOrder,
		}
	}

	return fieldsNode, nil
}
