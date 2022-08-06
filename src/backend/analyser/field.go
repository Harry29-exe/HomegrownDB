package analyser

import (
	"HomegrownDB/backend/analyser/anode"
	"HomegrownDB/backend/parser/pnode"
	"errors"
)

func AnalyseFields(node pnode.FieldsNode, tables anode.Tables) (anode.Fields, error) {
	fieldsCount := len(node.Fields)
	fieldsNode := anode.Fields{Fields: make([]anode.Field, fieldsCount)}

	for i, field := range node.Fields {
		table := tables.GetTableByAlias(field.TableAlias)
		if table == nil {
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
