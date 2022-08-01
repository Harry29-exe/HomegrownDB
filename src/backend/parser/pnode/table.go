package pnode

type TablesNode struct {
	Tables []TableNode
}

type TableNode struct {
	TableName  string
	TableAlias string
}
