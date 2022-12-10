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

type alias struct {
	node
	AliasName string
}
