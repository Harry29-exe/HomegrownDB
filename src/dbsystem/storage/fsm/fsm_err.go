package fsm

import (
	"HomegrownDB/dbsystem/hglib"
)

var _ hglib.DBError = NoFreeSpace{}

type NoFreeSpace struct {
}

func (n NoFreeSpace) Error() string {
	//TODO implement me
	panic("implement me")
}

func (n NoFreeSpace) Area() hglib.Area {
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
