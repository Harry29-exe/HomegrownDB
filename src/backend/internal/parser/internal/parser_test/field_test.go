package parser_test_test

import (
	. "HomegrownDB/backend/internal/parser/internal/parser"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
	"testing"
)

func TestField_Parse_ShouldParse(t *testing.T) {
	//given
	sentences := []testSentence{
		{"t1.col1", 2},
		{"t1.col1, t2.col2", 2},
		{"t1.col1 FROM ", 2},
	}
	expectedResult := pnode.FieldNode{
		TableAlias: "t1",
		FieldName:  "col1",
		FieldAlias: "col1",
	}

	for _, sentence := range sentences {
		source := createTestTokenSource(sentence.str, t)
		v := validator.NewValidator(source)
		//when
		result, err := Field.Parse(source, v)

		//then
		CorrectSentenceParserTestIsSuccessful(
			t, source, sentence,
			err, expectedResult, result)
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
		v := validator.NewValidator(source)
		//when
		_, err := Field.Parse(source, v)

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
