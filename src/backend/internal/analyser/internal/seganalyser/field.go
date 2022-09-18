package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/parser/pnode"
	"errors"
)

var Fields = fields{}

type fields struct{}

func (f fields) Analyse(node pnode.FieldsNode, tables anode.Tables, ctx *internal.AnalyserCtx) (anode.SelectFields, error) {
	fieldsCount := len(node.Fields)
	fieldsNode := anode.SelectFields{Fields: make([]anode.SelectField, fieldsCount)}

	for i, field := range node.Fields {
		table := tables.TableByAlias(field.TableAlias)
		if table == nil {
			return anode.SelectFields{}, errors.New("") // todo better message
		}

		if colOrder, ok := table.ColumnId(field.FieldName); !ok {
			return anode.SelectFields{}, errors.New("") // todo better message
		} else {
			fieldsNode.Fields[i] = anode.SelectField{
				Table:      table,
				Column:     table.GetColumn(colOrder),
				FieldAlias: field.FieldName,
			}
		}
	}

	return fieldsNode, nil
}
