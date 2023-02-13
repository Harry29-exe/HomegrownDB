package parser_test

import (
	"HomegrownDB/backend/internal/parser/parse"
	"HomegrownDB/backend/internal/parser/validator"
	pnode2 "HomegrownDB/backend/internal/pnode"
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

	expectedTree := pnode2.NewInsertStmt()
	expectedTree.Columns = []pnode2.ResultTarget{
		pnode2.NewResultTarget("", pnode2.NewColumnRef("name", "")),
		pnode2.NewResultTarget("", pnode2.NewColumnRef("age", "")),
	}
	expectedTree.SrcNode = pnode2.NewSelectStmtWithValues([][]pnode2.Node{
		{pnode2.NewAConstStr("bob"), pnode2.NewAConstInt(15)},
		{pnode2.NewAConstStr("Alice"), pnode2.NewAConstInt(24)},
	})
	expectedTree.Relation = pnode2.NewRangeVar("users", "")

	for _, query := range queries {
		//when
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)
		node, err := parse.Insert.Parse(source, v)

		//then
		assert.ErrIsNil(err, t)
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

	expectedTree := pnode2.NewInsertStmt()
	expectedTree.Columns = []pnode2.ResultTarget{
		pnode2.NewAStarResultTarget(),
	}
	expectedTree.SrcNode = pnode2.NewSelectStmtWithValues([][]pnode2.Node{
		{pnode2.NewAConstStr("bob"), pnode2.NewAConstInt(15)},
		{pnode2.NewAConstStr("Alice"), pnode2.NewAConstInt(24)},
	})
	expectedTree.Relation = pnode2.NewRangeVar("users", "")

	for _, query := range queries {
		//when
		source := newTestTokenSource(query)
		v := validator.NewValidator(source)
		node, err := parse.Insert.Parse(source, v)

		//then
		assert.ErrIsNil(err, t)
		assert.True(node.Equal(expectedTree), t)
		assert.Eq(len(source.Checkpoints), 0, t)
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

		_, err := parse.Insert.Parse(source, v)
		assert.ErrNotNil(err, t)
		assert.Eq(len(source.Checkpoints), 0, t)
	}
}
