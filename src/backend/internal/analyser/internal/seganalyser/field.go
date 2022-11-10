package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/tx"
	"errors"
)

var Fields = fields{}

type fields struct{}

func (f fields) Analyse(fieldNodes []pnode.FieldNode, tables anode.Tables, ctx *tx.Ctx) (anode.SelectFields, error) {
	fieldsCount := len(fieldNodes)
	fieldsNode := make([]anode.SelectField, fieldsCount)

	for i, field := range fieldNodes {
		table := tables.TableByAlias(field.TableAlias)
		if table == nil {
			return anode.SelectFields{}, errors.New("") // todo better message
		}

		if colOrder, ok := table.ColumnOrder(field.FieldName); !ok {
			return anode.SelectFields{}, errors.New("") // todo better message
		} else {
			qTableId, ok := ctx.CurrentQuery.GetQTableId(field.TableAlias)
			if !ok {
				return fieldsNode, errors.New("no table with alias: " + field.TableAlias)
			}

			fieldsNode[i] = anode.SelectField{
				Table:  qTableId,
				Column: colOrder,
			}
		}
	}

	return fieldsNode, nil
}
