package tokenizer_test

import (
	"HomegrownDB/backend/parser/tokenizer"
	"HomegrownDB/backend/parser/tokenizer/token"
	"strconv"
	"testing"
)

func TestTokenizerWithBasicSelectQuery(t *testing.T) {
	query := "sElECt  table_alias.column_name1 \n" +
		"FROM table_name table_alias \n" +
		"WHERE table_alias.column_name1 = 'str_value' AND \n" +
		"table_alias.column_name2 = 23  ;"

	tester := tokenizerTester{
		t:             t,
		testTokenizer: tokenizer.NewTokenizer(query),
	}

	tester.assertNextBasicToken(token.Select, "sElECt")
	tester.assertNextBasicToken(token.SpaceBreak, " ")
	tester.assertNextBasicToken(token.Identifier, "table_alias")
	tester.assertNextBasicToken(token.Dot, ".")
	//todo finish testing
}

type tokenizerTester struct {
	t             *testing.T
	testTokenizer tokenizer.Tokenizer
}

func (tt *tokenizerTester) passTokenizerError(err error) {
	if err != nil {
		tt.t.Error("Tokenizer returned following error:", err.Error())
	}
}

func (tt tokenizerTester) assertNextBasicToken(code token.Code, value string) {
	token, err := tt.testTokenizer.Next()
	tt.passTokenizerError(err)
	switch {
	case token.Value() != value:
		tt.t.Error("Tokenizer contains invalid token code: expected: \""+value+"\"",
			"actual: \""+token.Value()+"\"")
	case token.Code() != code:
		tt.t.Error("Token should has: "+strconv.Itoa(int(code))+
			"  code instead has ", strconv.Itoa(int(token.Code())))
	}
}
