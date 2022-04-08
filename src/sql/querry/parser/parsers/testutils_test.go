package parsers_test

import (
	"HomegrownDB/sql/querry/tokenizer"
	"HomegrownDB/sql/querry/tokenizer/token"
	"testing"
)

type testSentence struct {
	str        string // sentence for tokenizer
	pointerPos uint16 // expected pointer position after parsing
}

func newTestTokenSource(tokens []token.Token) *testTokenSource {
	return &testTokenSource{
		tokens:      tokens,
		tokensLen:   uint16(len(tokens)),
		pointer:     0,
		checkpoints: make([]uint16, 0, 5),
	}
}

type parseResult[T any] struct {
	node       T
	pointerPos uint16
}

type testTokenSource struct {
	tokens    []token.Token
	tokensLen uint16
	pointer   uint16

	checkpoints []uint16
}

func testSource(str string, t *testing.T) *testTokenSource {
	tknz := tokenizer.NewTokenizer(str)
	tokens := make([]token.Token, 0, 20)
	for tknz.HasNext() {
		newToken, err := tknz.Next()
		if err != nil {
			t.Error("tokenizer returned error during Fields parser test")
			t.FailNow()
		}
		tokens = append(tokens, newToken)
	}

	return newTestTokenSource(tokens)
}

func (t *testTokenSource) Next() token.Token {
	t.pointer++
	if t.pointer < t.tokensLen {
		return t.tokens[t.pointer]
	}
	t.pointer--
	return token.NilToken()
}

func (t *testTokenSource) Prev() token.Token {
	if t.pointer > 0 {
		t.pointer--
		return t.tokens[t.pointer]
	}

	return token.NilToken()
}

func (t *testTokenSource) Current() token.Token {
	return t.tokens[t.pointer]
}

func (t *testTokenSource) History() []token.Token {
	return t.tokens[0 : t.pointer+1]
}

func (t *testTokenSource) Checkpoint() {
	t.checkpoints = append(t.checkpoints, t.pointer)
}

func (t *testTokenSource) Commit() {
	lastIndex := len(t.checkpoints) - 1
	t.checkpoints = t.checkpoints[0:lastIndex]
}

func (t *testTokenSource) Rollback() {
	lastIndex := len(t.checkpoints) - 1
	t.pointer = t.checkpoints[lastIndex]
	t.checkpoints = t.checkpoints[0:lastIndex]
}

var ParserErr = parserError{}

type parserError struct{}

// OutputDiffers uses testing.T Error for printing error information,
// and marks test as filed with testing.T Fail
func (p parserError) OutputDiffers(t *testing.T, expected, output any, sentence string) {
	t.Error("Received output is different from expected one. "+
		"Expected: ", expected, "actual: ", output,
		". In sentence:\n", sentence)
	t.Fail()
}

func (p parserError) PointerPosDiffers(t *testing.T, expected, actual uint16, sentence string) {
	t.Error("TokenSource pointer position is different than"+
		"expected. Expected: ", expected, " actual: ", actual,
		".In sentence:\n", sentence)
}
