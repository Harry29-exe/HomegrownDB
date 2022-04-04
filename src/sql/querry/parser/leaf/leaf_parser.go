package leaf

import "HomegrownDB/sql/querry/tokenizer"

type Context interface {
	Tokenizer() tokenizer.Tokenizer
	NextToken() tokenizer.Token
}
