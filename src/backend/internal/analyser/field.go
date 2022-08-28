package analyser

import (
	anode2 "HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
	"errors"
)

func AnalyseFields(node pnode.FieldsNode, tables anode2.Tables) (anode2.Fields, error) {
	fieldsCount := len(node.Fields)
	fieldsNode := anode2.Fields{Fields: make([]anode2.Field, fieldsCount)}

	for i, field := range node.Fields {
		table := tables.GetTableByAlias(field.TableAlias)
		if table == nil {
			return anode2.Fields{}, errors.New("") // todo better message
		}

		if colOrder, ok := table.ColumnId(field.FieldName); !ok {
			return anode2.Fields{}, errors.New("") // todo better message
		} else {
			fieldsNode.Fields[i] = anode2.Field{
				Table:     table,
				Column:    table.GetColumn(colOrder),
				FieldName: field.FieldName,
			}
		}
	}

	return fieldsNode, nil
}
