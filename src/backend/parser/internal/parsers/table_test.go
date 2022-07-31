package parsers_test

import (
	"HomegrownDB/backend/parser/internal/parsers"
	"testing"
)

func TestTable_Parse_ShouldParse(t *testing.T) {
	//given
	sentences := []testSentence{
		{"table1 t1, table2 t2", 2},
		{"table1 t1, table2 t2 WHERE", 2},
		{"table1 t1", 2},
		{"table1 t1 ", 2},
	}

	expectedNode := parsers.TableNode{
		TableName:  "table1",
		TableAlias: "t1",
	}

	for _, sentence := range sentences {
		_tableParserPositiveTest(t, sentence, expectedNode)
	}
}

func TestTable_Parse_ShouldParse2(t *testing.T) {
	//given
	sentences := []testSentence{
		{"table1 WHERE ", 0},
		{"table1", 0},
		{"table1, table2 t2", 0},
	}

	expectedNode := parsers.TableNode{
		TableName:  "table1",
		TableAlias: "table1",
	}

	for _, sentence := range sentences {
		_tableParserPositiveTest(t, sentence, expectedNode)
	}
}
func _tableParserPositiveTest(t *testing.T, sentence testSentence, expectedNode parsers.TableNode) {
	source := createTestTokenSource(sentence.str, t)
	output, err := parsers.Table.Parse(source)

	CorrectSentenceParserTestIsSuccessful(
		t, source, sentence,
		err, expectedNode, *output)
}
