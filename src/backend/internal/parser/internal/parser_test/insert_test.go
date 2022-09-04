package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/parser"
	"HomegrownDB/common/tests"
	"testing"
)

func TestSimpleInsertParse(t *testing.T) {
	query := "INSERT INTO users (name, age) VALUES ('bob', 15), ('Alice', 24)"
	source := internal.NewTokenSource(query)

	node, err := parser.InsertParser.Parse(source)
	if err != nil {
		t.Error(err.Error())
	}

	tests.AssertEq(node.Table.TableName, "users", t)
	tests.AssertEq(node.Table.TableAlias, "users", t)

	columns := node.Columns.ColNames
	tests.AssertEq(len(columns), 2, t)
	tests.AssertEq(columns[0], "name", t)
	tests.AssertEq(columns[1], "age", t)

	values := node.Rows
	tests.AssertEq(len(values), 2, t)

	//todo finish writing test
	//val1 := values[0].Values
	//tests.AssertEq(len(val1), 2, t)
	//tests.AssertEq(val1[0], "bob", t)
	//tests.AssertEq(val1[1], 15, t)
	//val2 := values[0].Values
	//tests.AssertEq(len(val2), 2, t)
	//tests.AssertEq(val2[0], "Alice", t)
	//tests.AssertEq(val2[1], 24, t)

}
