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
