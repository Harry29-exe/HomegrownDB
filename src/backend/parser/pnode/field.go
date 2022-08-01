package pnode

type FieldsNode struct {
	Fields []FieldNode
}

func (f *FieldsNode) AddField(field FieldNode) {
	f.Fields = append(f.Fields, field)
}

type FieldNode struct {
	TableAlias string
	FieldName  string
	FieldAlias string
}
