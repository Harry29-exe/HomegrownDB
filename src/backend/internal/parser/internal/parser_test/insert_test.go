package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/parser"
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
	for _, query := range queries {
		source := internal.NewTokenSource(query)

		node, err := parser.InsertParser.Parse(source)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Eq(node.Table.TableName, "users", t)
		assert.Eq(node.Table.TableAlias, "users", t)

		columns := node.Columns.ColNames
		assert.Eq(len(columns), 2, t)
		assert.Eq(columns[0], "name", t)
		assert.Eq(columns[1], "age", t)

		values := node.Rows
		assert.Eq(len(values), 2, t)

		val1 := values[0].Values
		assert.Eq(len(val1), 2, t)
		assert.Eq(val1[0].V().(string), "bob", t)
		assert.Eq(val1[1].V().(int), 15, t)
		val2 := values[1].Values
		assert.Eq(len(val2), 2, t)
		assert.Eq(val2[0].V().(string), "Alice", t)
		assert.Eq(val2[1].V().(int), 24, t)
	}
}

func TestInsertParseWithDefaultColumn(t *testing.T) {
	queries := []string{
		"INSERT INTO users VALUES ('bob', 15), ('Alice', 24)",
		"INSERT INTO users  VALUES ( 'bob' , 15), ( 'Alice' , 24)",
		"INSERT INTO users VALUES ('bob',15),('Alice',24)",
	}

	for _, query := range queries {
		source := internal.NewTokenSource(query)

		node, err := parser.InsertParser.Parse(source)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Eq(node.Table.TableName, "users", t)
		assert.Eq(node.Table.TableAlias, "users", t)

		columns := node.Columns.ColNames
		assert.Eq(len(columns), 0, t)

		values := node.Rows
		assert.Eq(len(values), 2, t)

		val1 := values[0].Values
		assert.Eq(len(val1), 2, t)
		assert.Eq(val1[0].V().(string), "bob", t)
		assert.Eq(val1[1].V().(int), 15, t)
		val2 := values[1].Values
		assert.Eq(len(val2), 2, t)
		assert.Eq(val2[0].V().(string), "Alice", t)
		assert.Eq(val2[1].V().(int), 24, t)
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
		source := internal.NewTokenSource(query)

		_, err := parser.InsertParser.Parse(source)
		assert.NotNil(err, t)
	}
}
