package fsm

import "HomegrownDB/dbsystem/schema/table"

type Store interface {
	GetFSM(id table.Id) FreeSpaceMap
	CreateFSM(definition table.RDefinition)
	DeleteFSM(id table.Id)
}
