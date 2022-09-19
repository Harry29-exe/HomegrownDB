package pnode

type FieldsNode struct {
	Node
	Fields []FieldNode
}

func (f *FieldsNode) AddField(field FieldNode) {
	f.Fields = append(f.Fields, field)
}

type FieldNode struct {
	Node
	TableAlias string
	FieldName  string
	FieldAlias string
}
