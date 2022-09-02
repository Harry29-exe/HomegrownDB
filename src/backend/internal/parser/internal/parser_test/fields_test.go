package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal/parser"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
	"testing"
)

func TestFields_Parse_ShouldParseString(t *testing.T) {
	//given
	correctSentences := []string{
		"u.name, u.surname , c.name",
		"u.name, \t u.surname, c.name",
		"u.name, u.surname\n,\t c.name",
	}
	expectedFields := []pnode.FieldNode{
		{TableAlias: "u", FieldName: "name", FieldAlias: "name"},
		{TableAlias: "u", FieldName: "surname", FieldAlias: "surname"},
		{TableAlias: "c", FieldName: "name", FieldAlias: "name"},
	}

	for _, sentence := range correctSentences {
		source := createTestTokenSource(sentence, t)
		v := validator.NewValidator(source)

		//when
		fieldsNode, err := parser.Fields.Parse(source, v)
		if err != nil {
			t.Error("Fields parser returned following error: ", err)
			t.FailNow()
		}

		//then
		if len(fieldsNode.Fields) != 3 {
			t.Error("Len(FieldsNode.Fields) should equal 3 instead was:",
				len(fieldsNode.Fields))
			t.FailNow()
		}
		for i, field := range fieldsNode.Fields {
			if field != expectedFields[i] {
				t.Error("at ", i, "expected value: ", expectedFields[i],
					"got:", field)
			}
		}
	}
}

func TestFields_Parse_ShouldReturnError(t *testing.T) {
	//given
	badSentences := []string{
		"u.name, u.surname,",
		"u.name, u.surname, ",
		"u.name, u..surname",
	}

	for _, sentence := range badSentences {
		source := createTestTokenSource(sentence, t)
		v := validator.NewValidator(source)

		//when
		_, err := parser.Fields.Parse(source, v)

		//then
		if err == nil {
			t.Error(
				"Fields parser did not returned error when "+
					"parsing invalid sequence:\n",
				sentence)
		}
	}
}
