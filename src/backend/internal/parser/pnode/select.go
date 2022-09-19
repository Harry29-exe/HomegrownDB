package pnode

func NewSelect() Select {
	return Select{
		Node:   Node{},
		Fields: make([]FieldNode, 0, 20),
		Tables: make([]TableNode, 0, 10),
	}
}

type Select struct {
	Node
	Fields []FieldNode
	Tables []TableNode
}

func (s *Select) AddField(field FieldNode) {
	s.Fields = append(s.Fields, field)
}

func (s *Select) AddTable(table TableNode) {
	s.Tables = append(s.Tables, table)
}
