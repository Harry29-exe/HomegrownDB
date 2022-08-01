package parser_test_test

import (
	. "HomegrownDB/backend/parser/internal/parsers"
	"HomegrownDB/backend/parser/pnode"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []testSentence{
		{"SELECT t1.col1 FROM table1 t1", 10},
	}
	expectedNode := pnode.SelectNode{
		Fields: pnode.FieldsNode{Fields: []pnode.FieldNode{
			{
				TableAlias: "t1",
				FieldName:  "col1",
				FieldAlias: "col1",
			},
		}},
		Tables: pnode.TablesNode{Tables: []pnode.TableNode{
			{"table1", "t1"},
		}},
	}

	for _, sentence := range sentences {
		source := createTestTokenSource(sentence.str, t)
		selectNode, err := Select.Parse(source)

		CorrectSentenceParserTestIsSuccessful(
			t, source, sentence,
			err, expectedNode, *selectNode)
	}

}
