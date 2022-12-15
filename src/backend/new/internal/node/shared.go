package node

import (
	"fmt"
)

// -------------------------
//      Alias
// -------------------------

func NewAlias(aliasName string) Alias {
	return &alias{
		node:      node{tag: TagAlias},
		AliasName: aliasName,
	}
}

type Alias = *alias

var _ Node = &alias{}

type alias struct {
	node
	AliasName string
}

func (a Alias) dEqual(node Node) bool {
	raw := node.(Alias)
	return a.AliasName == raw.AliasName
}

func (a Alias) DPrint(nesting int) string {
	return fmt.Sprintf("%s{AliasName: %s}",
		a.dTag(nesting), a.AliasName)
}
