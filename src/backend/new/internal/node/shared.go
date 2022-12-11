package node

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

func (a Alias) DEqual(node Node) bool {
	if res, ok := nodeEqual(a, node); ok {
		return res
	}
	raw := node.(Alias)
	return a.AliasName == raw.AliasName
}
