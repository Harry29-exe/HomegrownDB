package plan

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

type Conditions = []Condition

type Condition struct {
	Values []ConditionValue
	Type   conditionType
}

type ConditionValue interface {
	Value() any //todo change any to ConditionRawValue somehow
}

type ConditionRawValue interface {
	ConditionTupleValue | int | uint | string
}

type ConditionTupleValue struct {
	column column.Def
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
