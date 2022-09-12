package internal

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/schema/table"
	"errors"
)

var Fields = fields{}

type fields struct{}

func (f fields) Analyse(node pnode.FieldsNode, tables anode.Tables, ctx *analyser.Ctx) (anode.Fields, error) {
	fieldsCount := len(node.Fields)
	fieldsNode := anode.Fields{Fields: make([]anode.Field, fieldsCount)}

	for i, field := range node.Fields {
		tableId := tables.TableIdByAlias(field.TableAlias)
		if tableId == 0 {
			return anode.Fields{}, errors.New("") // todo better message
		}

		if colOrder, ok := table.ColumnId(field.FieldName); !ok {
			return anode.Fields{}, errors.New("") // todo better message
		} else {
			fieldsNode.Fields[i] = anode.Field{
				Table:     table,
				Column:    table.GetColumn(colOrder),
				FieldName: field.FieldName,
			}
		}
	}

	return fieldsNode, nil
}
