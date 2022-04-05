package parser

import (
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer"
)

func Parse(query string) ptree.NodeType {
	tk := tokenizer.NewTokenizer(query)

}
