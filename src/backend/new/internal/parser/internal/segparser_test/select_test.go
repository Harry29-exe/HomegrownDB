package parser_test

import (
	"HomegrownDB/backend/new/internal/parser/internal/segparser"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []string{
		"SELECT t1.col1 FROM ttable1 t1",
	}

	rangeVar := pnode.NewRangeVar("ttable1", "t1")
	from := []pnode.RangeVar{rangeVar}

	col1Ref := pnode.NewColumnRef("col1", "t1")
	col1 := pnode.NewResultTarget("", &col1Ref)
	targetList := []pnode.ResultTarget{col1}

	for _, sentence := range sentences {
		source := newTestTokenSource(sentence)
		v := validator.NewValidator(source)
		selectNode, err := segparser.Select.Parse(source, v)

		if err != nil {
			t.Error(err)
		}

		assert.EqDeep(selectNode.Targets, targetList, t)
		assert.EqDeep(selectNode.From, from, t)
	}

}
