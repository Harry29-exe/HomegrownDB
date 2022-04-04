package partial

import (
	"HomegrownDB/sql/querry/parser/parsetree"
)

type Context interface {
}

type Parser interface {
	parse(ctx *Context) (parsetree.Node, error)
}
