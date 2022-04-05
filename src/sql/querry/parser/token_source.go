package parser

import (
	"HomegrownDB/sql/querry/parser/defs"
	tk "HomegrownDB/sql/querry/tokenizer"
)

func NewTokenSource(query string) defs.TokenSource {
	return &tokenSource{
		tokenCache:  make([]tk.Token, 0, 10),
		currentLen:  0,
		pointer:     0,
		tokenizer:   tk.NewTokenizer(query),
		checkpoints: make([]uint16, 0, 8),
	}
}

type tokenSource struct {
	tokenCache []tk.Token
	currentLen uint16
	pointer    uint16

	tokenizer tk.Tokenizer

	checkpoints []uint16
}

func (t *tokenSource) Next() tk.Token {
	t.pointer++
	if t.pointer < t.currentLen {
		return t.tokenCache[t.pointer]
	}

	if t.tokenizer.HasNext() {
		next, err := t.tokenizer.Next()
		if err != nil {
			panic("tokenizer returned error: " + err.Error())
		}

		t.tokenCache = append(t.tokenCache, next)
		return next
	} else {
		t.pointer--
		return nil
	}
}

func (t *tokenSource) Prev() tk.Token {
	if t.pointer < 0 {
		return nil
	}

	t.pointer--
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) Current() tk.Token {
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) History() []tk.Token {
	return t.tokenCache[0 : t.pointer+1]
}

func (t *tokenSource) Checkpoint() {
	t.checkpoints = append(t.checkpoints, t.pointer)
}

func (t *tokenSource) Commit() {
	lastIndex := len(t.checkpoints) - 1
	t.checkpoints = t.checkpoints[0:lastIndex]
}

func (t *tokenSource) Rollback() {
	lastIndex := len(t.checkpoints) - 1
	t.pointer = t.checkpoints[lastIndex]
	t.checkpoints = t.checkpoints[0:lastIndex]
}
