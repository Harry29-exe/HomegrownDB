package hglib

type genericErr struct {
	Msg string
}

func (g genericErr) Error() string {
	return g.Msg
}

func (g genericErr) Area() Area {
	return DBSystem
}

func (g genericErr) MsgCanBeReturnedToClient() bool {
	return false
}

// -------------------------
//      Errors
// -------------------------

var IntegerOverflowErr = genericErr{Msg: "integer overflow, operation not possible"}
