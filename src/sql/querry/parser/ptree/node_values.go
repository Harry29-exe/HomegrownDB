package ptree

type FieldNodeValue struct {
	TableAlias string
	FieldName  string
	FieldAlias string
}

type TableNodeValue struct {
	TableName  string
	TableAlias string
}

type TableJoinNodeValue struct {
	NewTable        TableNodeValue // NewTable table that will be joined to query with this join
	NewTableColName string         // NewTableColAlias table of NewTable on which join will be performed
	TableAlias      string         // TableAlias table that NewTable will be joined on
	TableColName    string         // TableColName column of TableAlias that join will be performed on
}

type CondValueNodeValue struct {
	ValueType CondValueType
	// Value if ValueType == ColumnIdentifier then it is ColumnIdentifierNodeValue,
	// otherwise it has one of these types: int, uint, string appropriately to ValueType
	Value any
}

type CondValueType = uint16

const (
	ColumnIdentifier CondValueType = iota
	Int
	Float
	String
)

type ColumnIdentifierNodeValue struct {
	TableAlias string
	ColumnName string
}
