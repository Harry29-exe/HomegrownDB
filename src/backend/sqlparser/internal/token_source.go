package internal

import (
	"HomegrownDB/backend/tokenizer"
	"HomegrownDB/backend/tokenizer/token"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If method reach array's end it returns token with code token.Nil without changing pointer
type TokenSource interface {
	Next() token.Token      // Next move pointer forwards and returns, if source has no more tokens it returns nil
	Prev() token.Token      // Prev move pointer backwards and returns, if source has no more tokens it returns nil
	Current() token.Token   // Current returns token which has been returned with last method
	History() []token.Token // History returns all token from beginning to the one that Next would return

	Checkpoint() // Checkpoint creates new checkpoint for parser to rollback
	// CommitAndCheckpoint todo probably because of new aproach to helpers this hould be deleted
	CommitAndCheckpoint() // CommitAndCheckpoint commits and creates new checkpoint
	Commit()              // Commit deletes last checkpoint
	Rollback()            // Rollback to last checkpoint and removes this checkpoint
}

func NewTokenSource(query string) TokenSource {
	return &tokenSource{
		tokenCache:  make([]token.Token, 0, 10),
		currentLen:  0,
		pointer:     0,
		tokenizer:   tokenizer.NewTokenizer(query),
		checkpoints: make([]uint16, 0, 8),
	}
}

type tokenSource struct {
	tokenCache []token.Token
	currentLen uint16
	pointer    uint16

	tokenizer tokenizer.Tokenizer

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
