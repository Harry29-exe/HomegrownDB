package stores

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/schema/table"
)

type Tables interface {
	WTablesDefs
	TablesLocks
	TablesIOs
}

type RTablesDefs interface {
	GetTable(name string) (table.Definition, error)
	Table(id table.Id) table.Definition
}

type WTablesDefs interface {
	RTablesDefs
	AllTables() []table.Definition
	AddTable(table table.WDefinition) error
	RemoveTable(id table.Id) error
}

type TablesLocks interface {
	WLockTable(table.Id)
	RLockTable(table.Id)

	WUnlockTable(table.Id)
	RUnlockTable(table.Id)
}

type TablesIOs interface {
	GetTableIO(name string) (access.TableDataIO, error)
	TableIO(id table.Id) access.TableDataIO
}
