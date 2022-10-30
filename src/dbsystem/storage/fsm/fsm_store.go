package fsm

import "HomegrownDB/dbsystem/schema/table"

type Store interface {
	GetFSM(id table.Id) FreeSpaceMap
	CreateFSM(definition table.Definition)
	DeleteFSM(id table.Id)
}
