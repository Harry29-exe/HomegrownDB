package stores

import (
	"HomegrownDB/dbsystem/io"
	"HomegrownDB/dbsystem/schema/table"
)

type Tables interface {
	GetTable(name string) (table.Definition, error)
	Table(id table.Id) table.Definition
	AllTables() []table.Definition
	AddTable(table table.WDefinition) error
	RemoveTable(id table.Id) error

	GetTableIO(name string) (io.TableDataIO, error)
	TableIO(id table.Id) io.TableDataIO
}
