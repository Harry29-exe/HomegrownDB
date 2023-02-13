package parser_test

//
//import (
//	"HomegrownDB/backend/internal/parser/internal/parse"
//	"HomegrownDB/backend/internal/parser/internal/validator"
//	"HomegrownDB/backend/internal/parser/pnode"
//	"testing"
//)
//
//func TestTable_Parse_ShouldParse(t *testing.T) {
//	//given
//	sentences := []testSentence{
//		{"ttable1 t1, table2 t2", 2},
//		{"ttable1 t1, table2 t2 WHERE", 2},
//		{"ttable1 t1", 2},
//		{"ttable1 t1 ", 2},
//	}
//
//	expectedNode := pnode.TableNode{
//		TableName:  "ttable1",
//		TableAlias: "t1",
//	}
//
//	for _, sentence := range sentences {
//		_tableParserPositiveTest(t, sentence, expectedNode)
//	}
//}
//
//func TestTable_Parse_ShouldParse2(t *testing.T) {
//	//given
//	sentences := []testSentence{
//		{"ttable1 WHERE ", 0},
//		{"ttable1", 0},
//		{"ttable1, table2 t2", 0},
//	}
//
//	expectedNode := pnode.TableNode{
//		TableName:  "ttable1",
//		TableAlias: "ttable1",
//	}
//
//	for _, sentence := range sentences {
//		_tableParserPositiveTest(t, sentence, expectedNode)
//	}
//}
//func _tableParserPositiveTest(t *testing.T, sentence testSentence, expectedNode pnode.TableNode) {
//	source := createTokenSourceAndTestIt(sentence.str, t)
//	v := validator.NewValidator(source)
//	output, err := parse.Table.Parse(source, v)
//
//	CorrectSentenceParserTestIsSuccessful(
//		t, source, sentence,
//		err, expectedNode, output)
//}
