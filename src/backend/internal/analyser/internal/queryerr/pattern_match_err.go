package queryerr

import (
	"HomegrownDB/backend/dberr"
	"fmt"
)

func NewPatternMatchError(recType, recVal, expectedType string, recreatedQuery string) PatternMatchError {
	return PatternMatchError{
		ReceivedType:   recType,
		ReceivedValue:  recVal,
		ExpectedType:   expectedType,
		RecreatedQuery: recreatedQuery,
	}
}

var (
	_ error         = PatternMatchError{}
	_ dberr.DBError = PatternMatchError{}
)

type PatternMatchError struct {
	ReceivedType   string
	ReceivedValue  string
	ExpectedType   string
	RecreatedQuery string
}

func (p PatternMatchError) Error() string {
	return fmt.Sprintf(
		"values does not match column pattern. Expected %s but recived %s with value %s.\n%s <-- Here",
		p.ExpectedType, p.ReceivedType, p.ReceivedValue, p.RecreatedQuery)
}

func (p PatternMatchError) Area() dberr.Area {
	return dberr.Analyser
}

func (p PatternMatchError) MsgCanBeReturnedToClient() bool {
	return true
}

func NewPatternMatchLenError(expectedLen, actualLen int, recreatedQuery string) PatternMatchLenError {
	return PatternMatchLenError{
		ExpectedValuesCount: expectedLen,
		ReceivedValuesCount: actualLen,
		RecreatedQuery:      recreatedQuery,
	}
}

var (
	_ error         = PatternMatchLenError{}
	_ dberr.DBError = PatternMatchLenError{}
)

type PatternMatchLenError struct {
	ExpectedValuesCount int
	ReceivedValuesCount int
	RecreatedQuery      string
}

func (p PatternMatchLenError) Error() string {
	return fmt.Sprintf(
		"value count: %d did not match column count: %d, in query:\n%s <-- here",
		p.ExpectedValuesCount, p.ReceivedValuesCount, p.RecreatedQuery,
	)
}

func (p PatternMatchLenError) Area() dberr.Area {
	return dberr.Analyser
}

func (p PatternMatchLenError) MsgCanBeReturnedToClient() bool {
	return true
}
