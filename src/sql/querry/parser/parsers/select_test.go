package parsers_test

import (
	. "HomegrownDB/sql/querry/parser/parsers"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []testSentence{
		{"SELECT t1.col1 FROM table1 t1", 10},
	}
	expectedNode := SelectNode{
		Fields: &FieldsNode{Fields: []*FieldNode{
			{
				TableAlias: "t1",
				FieldName:  "col1",
				FieldAlias: "col1",
			},
		}},
		Tables: &TablesNode{Tables: []*TableNode{
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
