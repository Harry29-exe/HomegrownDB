package parser_test_test

import (
	. "HomegrownDB/backend/internal/parser/internal/parser"
	pnode2 "HomegrownDB/backend/internal/parser/pnode"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []testSentence{
		{"SELECT t1.col1 FROM table1 t1", 10},
	}
	expectedNode := pnode2.SelectNode{
		Fields: pnode2.FieldsNode{Fields: []pnode2.FieldNode{
			{
				TableAlias: "t1",
				FieldName:  "col1",
				FieldAlias: "col1",
			},
		}},
		Tables: pnode2.TablesNode{Tables: []pnode2.TableNode{
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
