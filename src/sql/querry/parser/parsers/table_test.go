package parsers_test

import (
	"HomegrownDB/sql/querry/parser/parsers"
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
		source := testSource(sentence.str, t)
		output, err := parsers.Table.Parse(source)

		if err != nil {
			t.Error("Table parser returned error: ", err)
			t.Fail()
			continue
		}

		if *output != expectedNode {
			ParserErr.OutputDiffers(t, expectedNode, *output, sentence.str)
		}
		if sentence.pointerPos != source.pointer {
			ParserErr.PointerPosDiffers(t,
				sentence.pointerPos,
				source.pointer,
				sentence.str)
		}
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
		source := testSource(sentence.str, t)
		output, err := parsers.Table.Parse(source)

		if err != nil {
			t.Error("Table parser returned error: ", err)
			t.Fail()
			continue
		}

		if *output != expectedNode {
			ParserErr.OutputDiffers(t, expectedNode, *output, sentence.str)
		}
		if sentence.pointerPos != source.pointer {
			ParserErr.PointerPosDiffers(t,
				sentence.pointerPos,
				source.pointer,
				sentence.str)
		}
	}
}
