package plan

var _ Node = InsertNodeSrc{}

type InsertNodeSrc struct {
	Rows []InsertRowSrc
}

func (i InsertNodeSrc) Type() NodeType {
	//TODO implement me
	panic("implement me")
}

func (i InsertNodeSrc) Children() []Node {
	return nil
}

type InsertRowSrc struct {
	Fields []InsertFieldSrc
}

type InsertFieldSrc struct {
	// Src node that will provide field's value
	Src Node
	// Value actual value to insert
	Value []byte
}
