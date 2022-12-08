package node

// -------------------------
//      Alias
// -------------------------

func NewAlias(aliasName string) Alias {
	return &alias{
		node:      node{tag: TagAlias},
		aliasName: aliasName,
	}
}

type Alias = *alias

type alias struct {
	node
	aliasName string
}
