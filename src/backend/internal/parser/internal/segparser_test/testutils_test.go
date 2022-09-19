package parser_test

import (
	"HomegrownDB/backend/internal/parser/internal/tokenizer"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/pnode"
	"reflect"
	"testing"
)

type testSentence struct {
	str        string // sentence for tokenizer
	pointerPos uint32 // expected Pointer position after parsing
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

	// test Pointer position
	if source.Pointer != sentence.pointerPos {
		ParserErr.PointerPosDiffers(t, sentence.pointerPos, source.Pointer, sentence.str)
		return false
	}

	//test uncommitted Checkpoints
	if len(source.Checkpoints) > 0 {
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

func createTokenSourceAndTestIt(str string, t *testing.T) *testTokenSource {
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
		TokenCache:  tokens,
		CurrentLen:  uint32(len(tokens)),
		Pointer:     0,
		Tokenizer:   tknz,
		Checkpoints: make([]uint32, 0, 5),
	}
}

func newTestTokenSource(query string) *testTokenSource {
	return &testTokenSource{
		TokenCache:  make([]token.Token, 0, 10),
		CurrentLen:  0,
		Pointer:     0,
		Tokenizer:   tokenizer.NewTokenizer(query),
		Checkpoints: make([]uint32, 0, 8),
	}
}

type testTokenSource struct {
	TokenCache []token.Token
	CurrentLen uint32
	Pointer    uint32

	Tokenizer tokenizer.Tokenizer

	Checkpoints []uint32
}

func (t *testTokenSource) Next() token.Token {
	t.Pointer++
	if t.Pointer < t.CurrentLen {
		return t.TokenCache[t.Pointer]
	}

	if t.Tokenizer.HasNext() {
		next, err := t.Tokenizer.Next()
		if err != nil {
			next = token.NewErrorToken(err.Error())
		}

		t.TokenCache = append(t.TokenCache, next)
		t.CurrentLen++
		return next
	} else {
		t.Pointer--
		return token.NilToken()
	}
}

func (t *testTokenSource) Prev() token.Token {
	if t.Pointer < 0 {
		return token.NilToken()
	}

	t.Pointer--
	return t.TokenCache[t.Pointer]
}

func (t *testTokenSource) Current() token.Token {
	if len(t.TokenCache) == 0 && t.Tokenizer.HasNext() {
		t.Pointer--
		return t.Next()
	}
	return t.TokenCache[t.Pointer]
}

func (t *testTokenSource) CurrentTokenIndex() uint32 {
	return t.Pointer
}

func (t *testTokenSource) History() []token.Token {
	return t.TokenCache[0 : t.Pointer+1]
}

func (t *testTokenSource) Checkpoint() {
	t.Checkpoints = append(t.Checkpoints, t.Pointer)
}

func (t *testTokenSource) Commit() {
	lastIndex := len(t.Checkpoints) - 1
	t.Checkpoints = t.Checkpoints[0:lastIndex]
}

func (t *testTokenSource) CommitAndInitNode(node *pnode.Node) {
	node.SetTokenIndexes(t.Checkpoints[len(t.Checkpoints)-1], t.Pointer)
	t.Commit()
}

func (t *testTokenSource) Rollback() {
	lastIndex := len(t.Checkpoints) - 1
	t.Pointer = t.Checkpoints[lastIndex]
	t.Checkpoints = t.Checkpoints[0:lastIndex]
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

func (p parserError) PointerPosDiffers(t *testing.T, expected, actual uint32, sentence string) {
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
