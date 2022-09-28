package ctype

import (
	"HomegrownDB/dbsystem/dberr"
	"fmt"
)

func NewConvError(from, fromValue, to string) dberr.DBError {
	return ConvError{
		From:      from,
		FromValue: fromValue,
		To:        to,
	}
}

func NewInputConvErr(from InputType, fromValue string, to Type) dberr.DBError {
	return ConvError{
		From:      InputTypeToStr(from),
		FromValue: fromValue,
		To:        CTypeToStr(to),
	}
}

type ConvError struct {
	From      string
	FromValue string
	To        string
}

func (c ConvError) Error() string {
	return fmt.Sprintf("can not convert type: %s (%s) to type: %s",
		c.From, c.FromValue, c.To)
}

func (c ConvError) Area() dberr.Area {
	return dberr.Analyser
}

func (c ConvError) MsgCanBeReturnedToClient() bool {
	return true
}
