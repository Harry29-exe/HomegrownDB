package plan

type InsertRowSrc struct {
	Fields []InsertFieldSrc
}

type InsertFieldSrc = *insertFieldSrc

type insertFieldSrc struct {
	// Src node that will provide field's value
	Src Node
	// Value actual value to insert
	Value []byte
}
