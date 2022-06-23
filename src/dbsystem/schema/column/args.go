package column

import (
	"errors"
)

func ArgsBuilder() *argsBuilder {
	return &argsBuilder{argsMap: &argsMap{}}
}

func (b *argsBuilder) Type(columnType Type) *argsBuilder {
	b.argsMap.SetType(columnType)
	return b
}

func (b *argsBuilder) Name(name string) *argsBuilder {
	b.argsMap.SetName(name)
	return b
}

func (b *argsBuilder) Nullable(nullable bool) *argsBuilder {
	b.argsMap.SetNullable(nullable)
	return b
}

func (b *argsBuilder) Length(length uint32) *argsBuilder {
	b.argsMap.SetLength(length)
	return b
}

func (b *argsBuilder) Build() Args {
	return b.argsMap
}

type argsBuilder struct {
	argsMap *argsMap
}

type argsMap struct {
	columnType    Type
	columnTypeSet bool

	name    string
	nameSet bool

	nullable    bool
	nullableSet bool

	length    uint32
	lengthSet bool
}

func (am *argsMap) Type() (Type, error) {
	if !am.columnTypeSet {
		return "", errors.New("arg type not set")
	}
	return am.columnType, nil
}

func (am *argsMap) SetType(colType Type) {
	am.columnType = colType
	am.columnTypeSet = true
}

func (am *argsMap) Name() (string, error) {
	if !am.nameSet {
		return "", errors.New("arg name not set")
	}
	return am.name, nil
}

func (am *argsMap) SetName(name string) {
	am.name = name
	am.nameSet = true
}

func (am *argsMap) Nullable() (bool, error) {
	if !am.nullableSet {
		return false, errors.New("arg nullable not set")
	}
	return am.nullable, nil
}

func (am *argsMap) SetNullable(nullable bool) {
	am.nullable = nullable
	am.nullableSet = true
}

func (am *argsMap) Length() (uint32, error) {
	if !am.lengthSet {
		return am.length, errors.New("arg length not set")
	}
	return am.length, nil
}

func (am *argsMap) SetLength(length uint32) {
	am.length = length
	am.lengthSet = true
}
