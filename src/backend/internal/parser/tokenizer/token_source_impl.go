package tokenizer

import (
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
)

func NewTokenSource(query string) TokenSource {
	return &tokenSource{
		tokenCache:  make([]token2.Token, 0, 50),
		currentLen:  0,
		pointer:     0,
		tokenizer:   NewTokenizer(query),
		checkpoints: make([]uint32, 0, 10),
	}
}

type tokenSource struct {
	tokenCache []token2.Token
	currentLen uint32
	pointer    uint32

	tokenizer Tokenizer

	checkpoints []uint32
}

func (t *tokenSource) Next() token2.Token {
	t.pointer++
	if t.pointer < t.currentLen {
		return t.tokenCache[t.pointer]
	}

	if t.tokenizer.HasNext() {
		next, err := t.tokenizer.Next()
		if err != nil {
			next = token2.NewErrorToken(err.Error())
		}

		t.tokenCache = append(t.tokenCache, next)
		t.currentLen++
		return next
	} else {
		t.pointer--
		return token2.NilToken()
	}
}

func (t *tokenSource) Prev() token2.Token {
	if t.pointer < 0 {
		return token2.NilToken()
	}

	t.pointer--
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) Current() token2.Token {
	if len(t.tokenCache) == 0 && t.tokenizer.HasNext() {
		t.pointer--
		return t.Next()
	}
	return t.tokenCache[t.pointer]
}

func (t *tokenSource) CurrentTokenIndex() uint32 {
	return t.pointer
}

func (t *tokenSource) History() []token2.Token {
	return t.tokenCache[0 : t.pointer+1]
}

func (t *tokenSource) Get(index uint) token2.Token {
	return t.tokenCache[index]
}

func (t *tokenSource) GetPtrRelative(index int) token2.Token {
	return t.tokenCache[int(t.pointer)+index]
}

func (t *tokenSource) Checkpoint() {
	t.checkpoints = append(t.checkpoints, t.pointer)
}

func (t *tokenSource) Commit() {
	lastIndex := len(t.checkpoints) - 1
	t.checkpoints = t.checkpoints[0:lastIndex]
}

func (t *tokenSource) CommitAndInitNode(node pnode.Node) {
	node.SetStartToken(uint(t.checkpoints[len(t.checkpoints)-1]))
	node.SetEndToken(uint(t.pointer))
	t.Commit()
}

func (t *tokenSource) Rollback() {
	lastIndex := len(t.checkpoints) - 1
	t.pointer = t.checkpoints[lastIndex]
	t.checkpoints = t.checkpoints[0:lastIndex]
}
