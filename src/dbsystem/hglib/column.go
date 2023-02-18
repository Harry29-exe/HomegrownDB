package hglib

func NewInvalidColumnArg(argName string, reason string) InvalidColumnArg {
	return InvalidColumnArg{
		ArgName: argName,
		Reason:  reason,
	}
}

var _ DBError = InvalidColumnArg{}

type InvalidColumnArg struct {
	ArgName string
	Reason  string
}

func (i InvalidColumnArg) Error() string {
	//TODO implement me
	panic("implement me")
}

func (i InvalidColumnArg) Area() Area {
	//TODO implement me
	panic("implement me")
}

func (i InvalidColumnArg) MsgCanBeReturnedToClient() bool {
	//TODO implement me
	panic("implement me")
}
