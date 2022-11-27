package parser_test

import (
	"HomegrownDB/backend/new/internal/parser/internal/segparser"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSimpleInsertParse(t *testing.T) {
	//given
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
	expectedTree.SrcNode = pnode.NewSelectStmtWithValues([][]pnode.Node{
		{pnode.NewAConstStr("bob"), pnode.NewAConstInt(15)},
		{pnode.NewAConstStr("Alice"), pnode.NewAConstInt(24)},
	})
	expectedTree.Relation = pnode.NewRangeVar("users", "")

	for _, query := range queries {
		//when
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)
		node, err := segparser.Insert.Parse(source, v)

		//then
		assert.IsNil(err, t)
		assert.True(node.Equal(expectedTree), t)
		assert.Eq(len(source.Checkpoints), 0, t)
	}
}

func TestInsertParseWithDefaultColumn(t *testing.T) {
	//given
	queries := []string{
		"INSERT INTO users VALUES ('bob', 15), ('Alice', 24)",
		"INSERT INTO users  VALUES ( 'bob' , 15), ( 'Alice' , 24)",
		"INSERT INTO users VALUES ('bob',15),('Alice',24)",
	}

	expectedTree := pnode.NewInsertStmt()
	expectedTree.Columns = []pnode.ResultTarget{
		pnode.NewAStarResultTarget(),
	}
	expectedTree.SrcNode = pnode.NewSelectStmtWithValues([][]pnode.Node{
		{pnode.NewAConstStr("bob"), pnode.NewAConstInt(15)},
		{pnode.NewAConstStr("Alice"), pnode.NewAConstInt(24)},
	})
	expectedTree.Relation = pnode.NewRangeVar("users", "")

	for _, query := range queries {
		//when
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)
		node, err := segparser.Insert.Parse(source, v)

		//then
		assert.Eq(len(source.Checkpoints), 0, t)
		assert.IsNil(err, t)
		assert.True(node.Equal(expectedTree), t)
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
