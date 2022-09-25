package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal/segparser"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []string{
		"SELECT t1.col1 FROM table1 t1",
	}

	tableAlias := "t1"
	tableName := "table1"
	fieldName := "col1"
	fieldAlias := "col1"

	for _, sentence := range sentences {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		if err != nil {
			t.Error(err)
		}

		assert.Eq(len(selectNode.Tables), 1, t)
		table := selectNode.Tables[0]
		assert.Eq(table.TableName, tableName, t)
		assert.Eq(table.TableAlias, tableAlias, t)

		assert.Eq(len(selectNode.Fields), 1, t)
		field := selectNode.Fields[0]
		assert.Eq(field.FieldName, fieldName, t)
		assert.Eq(field.FieldAlias, fieldAlias, t)
	}

}
