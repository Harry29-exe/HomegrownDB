package plan

var _ Node = InsertValuesSrc{}

type InsertValuesSrc struct {
	Rows []InsertRowSrc
}

func (i InsertValuesSrc) Type() NodeType {
	//TODO implement me
	panic("implement me")
}

func (i InsertValuesSrc) Children() []Node {
	return nil
}

type InsertRowSrc struct {
	Fields []InsertFieldSrc
}

type InsertFieldSrc struct {
	// Src test that will provide field's value
	Src Node
	// Value actual value to insert
	Value []byte
}
