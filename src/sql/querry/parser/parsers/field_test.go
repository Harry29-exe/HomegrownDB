package parsers_test

import (
	. "HomegrownDB/sql/querry/parser/parsers"
	"testing"
)

func TestField_Parse_ShouldParse(t *testing.T) {
	//given
	sentences, expectedResults := []string{
		"t1.col1",   //fully parsed
		"t2a2.col2", //fully parsed
		"t2.col2 FROM ",
	}, []parseResult[FieldNode]{
		{
			node:       FieldNode{TableAlias: "t1", FieldName: "col1", FieldAlias: "col1"},
			pointerPos: uint16(2),
		},
		{
			node:       FieldNode{TableAlias: "t2a2", FieldName: "col2", FieldAlias: "col2"},
			pointerPos: uint16(2),
		},
		{
			node:       FieldNode{TableAlias: "t2", FieldName: "col2", FieldAlias: "col2"},
			pointerPos: uint16(2),
		},
	}

	for i, sentence := range sentences {
		source := testSource(sentence, t)

		//when
		result, err := Field.Parse(source)

		//then
		if err != nil {
			t.Error("Field parser returned following error:\n",
				err,
				"when given valid sentence:\n\n",
				sentence)
			t.Fail()
		}

		expectedResult := expectedResults[i]
		if *result != expectedResult.node {
			t.Error("result: ", *result, " does not match expected value: ",
				expectedResults[i].node,
				"\nafter parsing sentence: ", sentence)
			t.Fail()
		}

		if expectedResult.pointerPos != source.pointer {
			t.Error("after parsing following sentence: ", sentence,
				"\npointer doesn't end up in expected position. Expected: ",
				expectedResult.pointerPos, " Real: ", source.pointer)
			t.Fail()
		}
	}
}

func TestField_Parse_ShouldReturnError(t *testing.T) {
	//given
	sentences := []string{
		"t1.2",
		"t1,col1",
		"t1 .col1",
	}
	expectedPointerPos := []uint16{
		0,
		0,
		0,
	}

	for i, sentence := range sentences {
		source := testSource(sentence, t)
		//when
		_, err := Field.Parse(source)

		if err == nil {
			t.Error("Field parser returned nil error for invalid token sequence: ",
				sentence)
			t.Fail()
		}

		if source.pointer != expectedPointerPos[i] {
			t.Error("after parsing following sentence: ", sentence,
				"\npointer doesn't end up in expected position. Expected: ",
				expectedPointerPos[i], " Real: ", source.pointer)
			t.Fail()
		}
	}
}
