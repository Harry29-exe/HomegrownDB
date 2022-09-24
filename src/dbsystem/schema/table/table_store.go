package table

type Store interface {
	WDefsStore
}

type RDefsStore interface {
	GetTable(name string) (Definition, error)
	Table(id Id) Definition
}

type WDefsStore interface {
	RDefsStore
	AllTables() []Definition
	AddTable(table WDefinition) error
	RemoveTable(id Id) error
}
