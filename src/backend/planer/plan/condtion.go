package plan

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type Conditions = []Condition

type Condition struct {
	LeftValue any
}

type ConditionValue interface {
	ValueType() conditionValuesType
	Value() ConditionRawValue
}

type ConditionRawValue interface {
	ConditionTupleValue | int | uint | string
}

type ConditionTupleValue struct {
	column column.Definition
	table  table.Definition
}

type conditionType = uint8

const (
	AND conditionType = iota
	OR
	EQ
	GEQ
	LEQ
	LIKE
)

type conditionValuesType = uint8

const (
	ColumnAndColumn conditionType = iota
	ColumnAndValue
	ValueAndValue
)
