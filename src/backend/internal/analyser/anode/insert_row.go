package anode

type InsertRow struct {
	Fields []InsertField
}

type InsertField struct {
	Value      []byte
	Expression any
}
