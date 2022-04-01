package tokenizer

import "strings"

type Tokenizer interface {
	Next() Token
	IsEmpty() bool
}

func NewTokenizer(str string) Tokenizer {
	return newBasicTokenizer(str)
}

type basicTokenizer struct {
	str string
	words
}

func newBasicTokenizer(str string) *basicTokenizer {

}

func (t *basicTokenizer) Next() Token {
	reader := strings.NewReader(t.str)
	reader.
}

func (t *basicTokenizer) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}
