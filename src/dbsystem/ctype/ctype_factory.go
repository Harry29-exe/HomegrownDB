package ctype

import (
	"HomegrownDB/dbsystem/dberr"
	"fmt"
)

func CreateCType(cType Type, args map[string]any) (CType, error) {
	switch cType {
	case TypeInt8:
		return int8Factory.Build(args)
	case TypeStr:
		return strFactory.Build(args)
	default:
		panic(fmt.Sprintf("not supported ctype: %s", CTypeToStr(cType)))
	}
}

type factory interface {
	Build(args map[string]any) (CType, dberr.DBError)
}

func NewFactoryArgError(cType, argName, expectedType string) FactoryArgError {
	return FactoryArgError{
		msg: fmt.Sprintf("can not create column with type %s, because argument: %s is not of type %s",
			cType, argName, expectedType),
	}
}

type FactoryArgError struct{ msg string }

func (f FactoryArgError) Error() string {
	return f.msg
}

func (f FactoryArgError) Area() dberr.Area {
	return dberr.DBSystem
}

func (f FactoryArgError) MsgCanBeReturnedToClient() bool {
	return true
}
