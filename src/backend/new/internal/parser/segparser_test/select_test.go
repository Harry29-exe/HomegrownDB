package parser_test

import (
	"HomegrownDB/backend/new/internal/parser/segparser"
	"HomegrownDB/backend/new/internal/parser/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []string{
		"SELECT t1.col1 FROM ttable1 t1",
	}

	expectedStmt := pnode.NewSelectStmt()
	expectedStmt.From = []pnode.Node{pnode.NewRangeVar("ttable1", "t1")}
	expectedStmt.Targets = []pnode.ResultTarget{
		pnode.NewResultTarget("", pnode.NewColumnRef("col1", "t1")),
	}

	for _, sentence := range sentences {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		assert.ErrIsNil(err, t)
		assert.True(selectNode.Equal(expectedStmt), t)
	}
}

func TestSelectParser_Parse_ShouldParse2(t *testing.T) {
	sentences := []string{
		"SELECT u.name, u.surname, u.age FROM users u",
		"SELECT u.name,u.surname,u.age FROM users u",
		"SELECT u.name , u.surname , u.age FROM users u",
		"SELECT u.name, u.surname, u.age\nFROM users u",
		"SELECT\nu.name,\nu.surname,\nu.age\nFROM\nusers\nu",
	}

	expectedStmt := pnode.NewSelectStmt()
	expectedStmt.From = []pnode.Node{pnode.NewRangeVar("users", "u")}
	expectedStmt.Targets = []pnode.ResultTarget{
		pnode.NewResultTarget("", pnode.NewColumnRef("name", "u")),
		pnode.NewResultTarget("", pnode.NewColumnRef("surname", "u")),
		pnode.NewResultTarget("", pnode.NewColumnRef("age", "u")),
	}

	for _, sentence := range sentences {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		assert.ErrIsNil(err, t)
		assert.True(selectNode.Equal(expectedStmt), t)
	}
}
