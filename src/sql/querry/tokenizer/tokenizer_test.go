package tokenizer_test

import (
	"HomegrownDB/sql/querry/tokenizer"
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

	tester.assertNextBasicToken(tokenizer.Select, "sElECt")
	tester.assertNextBasicToken(tokenizer.Text, "table_alias")
	tok := tester.testTokenizer
	for tok.HasNext() {
		next, _ := tok.Next()
		print("[", next.Code(), "-", next.Value(), "]")
	}
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

func (tt tokenizerTester) assertNextBasicToken(code tokenizer.TokenCode, value string) {
	token, err := tt.testTokenizer.Next()
	tt.passTokenizerError(err)
	switch {
	case token.Value() != value:
		tt.t.Error("Tokenizer contains invalid value for select. Should contain: \""+value+"\"",
			"contains: \""+token.Value()+"\"")
	case token.Code() != code:
		tt.t.Error("Token should has: "+strconv.Itoa(int(code))+
			"  code instead has ", strconv.Itoa(int(token.Code())))
	}
}
