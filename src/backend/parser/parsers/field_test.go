package parsers_test

import (
	. "HomegrownDB/backend/parser/parsers"
	. "HomegrownDB/sql/backend/parser/parsers"
	"testing"
)

func TestField_Parse_ShouldParse(t *testing.T) {
	//given
	sentences := []testSentence{
		{"t1.col1", 2},
		{"t1.col1, t2.col2", 2},
		{"t1.col1 FROM ", 2},
	}
	expectedResult := FieldNode{
		TableAlias: "t1",
		FieldName:  "col1",
		FieldAlias: "col1",
	}

	for _, sentence := range sentences {
		source := createTestTokenSource(sentence.str, t)

		//when
		result, err := Field.Parse(source)

		//then
		CorrectSentenceParserTestIsSuccessful(
			t, source, sentence,
			err, expectedResult, *result)
	}
}

func TestField_Parse_ShouldReturnError(t *testing.T) {
	//given
	sentences := []testSentence{
		{"t1.2", 0},
		{"t1,col1", 0},
		{"t1 .col1", 0},
	}

	for _, sentence := range sentences {
		source := createTestTokenSource(sentence.str, t)
		//when
		_, err := Field.Parse(source)

		if err == nil {
			t.Error("Field parser returned nil error for invalid token sequence: ",
				sentence)
			t.Fail()
		}

		if source.pointer != sentence.pointerPos {
			t.Error("after parsing following sentence: ", sentence,
				"\npointer doesn't end up in expected position. Expected: ",
				sentence.pointerPos, " Real: ", source.pointer)
			t.Fail()
		}
	}
}
