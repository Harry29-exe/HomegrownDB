package pnode

type TokenHolder interface {
	RecreateQueryToToken(tokenIndex TokenIndex) string
}

type TokenIndex = uint32

type Node struct {
	NodeStartTokenIndex TokenIndex
	NodeEndTokenIndex   TokenIndex
	TokenHolder         TokenHolder
}

func (n Node) RecreateQuery() string {
	return n.TokenHolder.RecreateQueryToToken(n.NodeEndTokenIndex + 1)
}
