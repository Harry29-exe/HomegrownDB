package queryerr

import (
	"HomegrownDB/backend/dberr"
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/parser/pnode"
)

func NewPatternMatchError(recType, recVal, expectedType string, node pnode.Node) PatternMatchError {
	return PatternMatchError{
		ReceivedType:  recType,
		ReceivedValue: recVal,
		ExpectedType:  expectedType,
		Node:          node,
	}
}

var (
	_ error         = PatternMatchError{}
	_ dberr.DBError = PatternMatchError{}
)

type PatternMatchError struct {
	ReceivedType  string
	ReceivedValue string
	ExpectedType  string
	Node          pnode.Node
}

func (p PatternMatchError) Error() string {
	return ""
}

func (p PatternMatchError) Area() dberr.Area {
	//TODO implement me
	panic("implement me")
}

func (p PatternMatchError) MsgCanBeReturnedToClient() bool {
	//TODO implement me
	panic("implement me")
}

func NewPatternMatchLenError(expectedLen, actualLen int, node pnode.Node, ctx internal.AnalyserCtx) PatternMatchLenError {
	//todo implement me
	panic("Not implemented")
}

type PatternMatchLenError struct {
	ExpectedValuesCount int
	ReceivedValuesCount int
	Node                pnode.Node
}

func (p PatternMatchLenError) Error() string {
	//TODO implement me
	panic("implement me")
}

func (p PatternMatchLenError) Area() dberr.Area {
	//TODO implement me
	panic("implement me")
}

func (p PatternMatchLenError) MsgCanBeReturnedToClient() bool {
	//TODO implement me
	panic("implement me")
}
