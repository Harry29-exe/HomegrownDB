package internal

import (
	tk "HomegrownDB/backend/internal/parser/internal/tokenizer"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

func NewTokenSource(query string) TokenSource {
	return &tokenSource{
		tokenCache:  make([]token.Token, 0, 10),
		currentLen:  0,
		pointer:     0,
		tokenizer:   tk.NewTokenizer(query),
		checkpoints: make([]uint16, 0, 8),
	}
}

type tokenSource struct {
	tokenCache []token.Token
	currentLen uint16
	pointer    uint16

	tokenizer tk.Tokenizer

	checkpoints []uint16
}

func (t *tokenSource) Next() token.Token {
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
		t.currentLen++
		return next
	} else {
		t.pointer--
		return token.NilToken()
	}
}

func (t *tokenSource) Prev() token.Token {
	if t.pointer < 0 {
		return token.NilToken()
	}

	t.pointer--
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) Current() token.Token {
	if len(t.tokenCache) == 0 && t.tokenizer.HasNext() {
		t.pointer--
		return t.Next()
	}
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) History() []token.Token {
	return t.tokenCache[0 : t.pointer+1]
}

func (t *tokenSource) Checkpoint() {
	t.checkpoints = append(t.checkpoints, t.pointer)
}

func (t *tokenSource) CommitAndCheckpoint() {
	t.checkpoints[len(t.checkpoints)-1] = t.pointer
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
