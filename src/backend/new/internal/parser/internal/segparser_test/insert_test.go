package parser_test

import (
	"HomegrownDB/backend/new/internal/parser/internal/segparser"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSimpleInsertParse(t *testing.T) {
	queries := []string{
		"INSERT INTO users (name, age) VALUES ('bob', 15), ('Alice', 24)",
		"INSERT INTO users  (name , age ) VALUES ( 'bob' , 15), ('Alice', 24) ",
		"INSERT INTO users (name,age) VALUES ('bob',15),('Alice' ,   24 )",
		"INSERT INTO users  (  name  ,  age  ) VALUES ('bob',15)  , (  'Alice'   ,   24   )",
	}

	expectedTree := pnode.NewInsertStmt()
	expectedTree.Columns = []pnode.ResultTarget{
		pnode.NewResultTarget("", pnode.NewColumnRef("name", "")),
		pnode.NewResultTarget("", pnode.NewColumnRef("age", "")),
	}

	for _, query := range queries {
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)

		node, err := segparser.Insert.Parse(source, v)
		assert.Eq(len(source.Checkpoints), 0, t)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Eq(node.Table.TableName, "users", t)
		assert.Eq(node.Table.TableAlias, "users", t)

		columns := node.ColNames
		assert.Eq(len(columns), 2, t)
		assert.Eq(columns[0], "name", t)
		assert.Eq(columns[1], "age", t)

		rows := node.Rows
		assert.Eq(len(rows), 2, t)

		row1 := rows[0].Fields
		assert.Eq(len(row1), 2, t)
		assert.Eq(row1[0].Value.V.(string), "bob", t)
		assert.Eq(row1[1].Value.V.(int64), 15, t)
		row2 := rows[1].Fields
		assert.Eq(len(row2), 2, t)
		assert.Eq(row2[0].Value.V.(string), "Alice", t)
		assert.Eq(row2[1].Value.V.(int64), 24, t)
	}
}

func TestInsertParseWithDefaultColumn(t *testing.T) {
	queries := []string{
		"INSERT INTO users VALUES ('bob', 15), ('Alice', 24)",
		"INSERT INTO users  VALUES ( 'bob' , 15), ( 'Alice' , 24)",
		"INSERT INTO users VALUES ('bob',15),('Alice',24)",
	}

	for _, query := range queries {
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)

		node, err := segparser.Insert.Parse(source, v)
		assert.Eq(len(source.Checkpoints), 0, t)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Eq(node.Table.TableName, "users", t)
		assert.Eq(node.Table.TableAlias, "users", t)

		columns := node.ColNames
		assert.Eq(len(columns), 0, t)

		values := node.Rows
		assert.Eq(len(values), 2, t)

		val1 := values[0].Fields
		assert.Eq(len(val1), 2, t)
		assert.Eq(val1[0].Value.V.(string), "bob", t)
		assert.Eq(val1[1].Value.V.(int64), 15, t)
		val2 := values[1].Fields
		assert.Eq(len(val2), 2, t)
		assert.Eq(val2[0].Value.V.(string), "Alice", t)
		assert.Eq(val2[1].Value.V.(int64), 24, t)
	}
}

func TestInsertParseInvalidQuery(t *testing.T) {
	queries := []string{ //todo check if returned errors are correct
		"INSERT INTO users VALUS ('bob', 15)",
		"INSERT INTO users VALUES ('bob, 15)",
		"INSERT INTO users VALUES 'bob', 15))",
		"INSERT INTO users VALUES ('bob', 15",
	}

	for _, query := range queries {
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)

		_, err := segparser.Insert.Parse(source, v)
		assert.NotNil(err, t)
		assert.Eq(len(source.Checkpoints), 0, t)
	}
}
