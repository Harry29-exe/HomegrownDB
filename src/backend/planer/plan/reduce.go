package plan

type SelectFields struct {
	Fields []SelectedField
	Child  Node
}

type SelectedField struct {
	Name    string
	FieldId FieldId
}

func (s SelectFields) Type() NodeType {
	return SelectFieldsNode
}

func (s SelectFields) Children() []Node {
	return []Node{s.Child}
}