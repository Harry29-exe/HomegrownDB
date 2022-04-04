package parser

import (
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
)

type context struct {
	tokenizer tokenizer.Tokenizer
	nextToken tokenizer.Token
	root      parsetree.Node
}

type partialParser interface {
	parse(ctx *context) error
}

func Parse(query string) parsetree.NodeType {

}
