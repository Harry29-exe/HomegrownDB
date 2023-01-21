package parser_test

import (
	"HomegrownDB/backend/internal/parser/segparser"
	"HomegrownDB/backend/internal/parser/validator"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	queryVariant := []string{
		"SELECT t1.col1 FROM ttable1 t1",
	}

	expectedStmt := pnode2.NewSelectStmt()
	expectedStmt.From = []pnode2.Node{pnode2.NewRangeVar("ttable1", "t1")}
	expectedStmt.Targets = []pnode2.ResultTarget{
		pnode2.NewResultTarget("", pnode2.NewColumnRef("col1", "t1")),
	}

	for _, sentence := range queryVariant {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		assert.ErrIsNil(err, t)
		assert.True(selectNode.Equal(expectedStmt), t)
	}
}

func TestSelectParser_Parse_ShouldParse2(t *testing.T) {
	queryVariant := []string{
		"SELECT u.name, u.surname, u.age FROM users u",
		"SELECT u.name,u.surname,u.age FROM users u",
		"SELECT   u.name   ,   u.surname   ,   u.age   FROM   users   u",
		"SELECT u.name, u.surname, u.age\nFROM users u",
		"SELECT u.name, u.surname, u.age FROM\nusers u",
		"SELECT\nu.name,\nu.surname,\nu.age\nFROM\nusers\nu",
	}

	expectedStmt := pnode2.NewSelectStmt()
	expectedStmt.From = []pnode2.Node{pnode2.NewRangeVar("users", "u")}
	expectedStmt.Targets = []pnode2.ResultTarget{
		pnode2.NewResultTarget("", pnode2.NewColumnRef("name", "u")),
		pnode2.NewResultTarget("", pnode2.NewColumnRef("surname", "u")),
		pnode2.NewResultTarget("", pnode2.NewColumnRef("age", "u")),
	}

	for _, sentence := range queryVariant {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		assert.ErrIsNil(err, t)
		assert.True(selectNode.Equal(expectedStmt), t)
	}
}

func TestSelectParser_Parse_ShouldParse3(t *testing.T) {
	sentences := []string{
		"SELECT u.name, u.surname, u.age FROM users u",
	}

	expectedStmt := pnode2.NewSelectStmt()
	expectedStmt.From = []pnode2.Node{pnode2.NewRangeVar("users", "u")}
	expectedStmt.Targets = []pnode2.ResultTarget{
		pnode2.NewResultTarget("", pnode2.NewColumnRef("name", "u")),
		pnode2.NewResultTarget("", pnode2.NewColumnRef("surname", "u")),
		pnode2.NewResultTarget("", pnode2.NewColumnRef("age", "u")),
	}

	for _, sentence := range sentences {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		assert.ErrIsNil(err, t)
		assert.True(selectNode.Equal(expectedStmt), t)
	}
}
