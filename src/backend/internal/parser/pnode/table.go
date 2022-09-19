package pnode

type TablesNode struct {
	Node
	Tables []TableNode
}

type TableNode struct {
	Node
	TableName  string
	TableAlias string
}
