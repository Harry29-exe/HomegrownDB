package fsm

import "HomegrownDB/dbsystem/dberr"

var _ dberr.DBError = NoFreeSpace{}

type NoFreeSpace struct {
}

func (n NoFreeSpace) Error() string {
	//TODO implement me
	panic("implement me")
}

func (n NoFreeSpace) Area() dberr.Area {
	//TODO implement me
	panic("implement me")
}

func (n NoFreeSpace) MsgCanBeReturnedToClient() bool {
	//TODO implement me
	panic("implement me")
}

type internalError = uint8

const (
	none internalError = iota
	corrupted
	noSpace
)
