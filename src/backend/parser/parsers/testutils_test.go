package parsers_test

import (
	"HomegrownDB/backend/tokenizer"
	"HomegrownDB/backend/tokenizer/token"
	"reflect"
	"testing"
)

type testSentence struct {
	str        string // sentence for tokenizer
	pointerPos uint16 // expected pointer position after parsing
}

func createTestTokenSource(str string, t *testing.T) *testTokenSource {
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

	return &testTokenSource{
		tokens:      tokens,
		tokensLen:   uint16(len(tokens)),
		pointer:     0,
		checkpoints: make([]uint16, 0, 5),
	}
}

func CorrectSentenceParserTestIsSuccessful(
	t *testing.T,
	source *testTokenSource,
	sentence testSentence,
	parseErr error,
	expectedNode any,
	actualNode any,
) bool {
	// test if error is nil
	if parseErr != nil {
		ParserErr.ParserReturnedErr(t, parseErr, sentence.str)
		return false
	}

	// test pointer position
	if source.pointer != sentence.pointerPos {
		ParserErr.PointerPosDiffers(t, sentence.pointerPos, source.pointer, sentence.str)
		return false
	}

	//test uncommitted checkpoints
	if len(source.checkpoints) > 0 {
		ParserErr.UncommittedCheckpoint(t, sentence.str)
		return false
	}

	// test nodes equals

	if !reflect.DeepEqual(expectedNode, actualNode) {
		ParserErr.OutputDiffers(t, expectedNode, actualNode, sentence.str)
		return false
	}

	return true
}

type testTokenSource struct {
	tokens    []token.Token
	tokensLen uint16
	pointer   uint16

	checkpoints []uint16
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

func (t *testTokenSource) CommitAndCheckpoint() {
	t.checkpoints[len(t.checkpoints)-1] = t.pointer
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
		"\nIn sentence:\""+sentence+"\"")
	t.Fail()
}

func (p parserError) PointerPosDiffers(t *testing.T, expected, actual uint16, sentence string) {
	t.Error("TokenSource pointer position is different than",
		"expected. Expected: ", expected, " actual: ", actual,
		"\nIn sentence:\""+sentence+"\"")
	t.Fail()
}

func (p parserError) ParserReturnedErr(t *testing.T, err error, sentence string) {
	t.Error("Parser returned unexpected error: ", err, " while parsing following sentence:\n\"", sentence, "\n")
	t.Fail()
}

func (p parserError) UncommittedCheckpoint(t *testing.T, sentence string) {
	t.Error("Parser left uncommitted checkpoint parsing following sentence:\n\"", sentence, "\"")
	t.Fail()
}
