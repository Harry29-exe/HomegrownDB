package table

type Store interface {
	DefsStore
}

type RDefsStore interface {
	// FindTable searches for table with given name
	// if it exists it returns table id otherwise it
	// returns InvalidTableId
	FindTable(name string) Id

	AccessTable(id Id, lockMode TableLockMode) Definition
}

type DefsStore interface {
	RDefsStore
	AddNewTable(table Definition) error
	LoadTable(table Definition) error
	RemoveTable(id Id) error
}

type TableLockMode = uint8

const (
	RLockMode TableLockMode = iota
	WLockMode
)
