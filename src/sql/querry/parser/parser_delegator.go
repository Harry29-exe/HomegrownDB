package parser

import "HomegrownDB/sql/querry/tokenizer"

type cache struct {
	tokenizer    tokenizer.Tokenizer
	currentToken tokenizer.Token
}

func Parse(query string) {

}
