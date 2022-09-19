package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal/segparser"
	"HomegrownDB/backend/internal/parser/pnode"
	"testing"
)

func TestSelectParser_Parse_ShouldParse(t *testing.T) {
	sentences := []testSentence{
		{"SELECT t1.col1 FROM table1 t1", 10},
	}
	expectedNode := pnode.Select{
		Fields: []pnode.FieldNode{
			{
				TableAlias: "t1",
				FieldName:  "col1",
				FieldAlias: "col1",
			},
		},
		Tables: []pnode.TableNode{
			{
				TableName:  "table1",
				TableAlias: "t1",
			},
		},
	}

	for _, sentence := range sentences {
		source := createTokenSourceAndTestIt(sentence.str, t)
		selectNode, err := segparser.Select.Parse(source)

		CorrectSentenceParserTestIsSuccessful(
			t, source, sentence,
			err, expectedNode, *selectNode)
	}

}
