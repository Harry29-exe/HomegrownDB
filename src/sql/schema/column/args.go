package column

import "errors"

type Args interface {
	Type() (Type, error)
	Name() (string, error)
	Nullable() (bool, error)
	Length() (uint32, error)
}

type ArgsMap struct {
	columnType    Type
	columnTypeSet bool

	name    string
	nameSet bool

	nullable    bool
	nullableSet bool

	length    uint32
	lengthSet bool
}

func (am *ArgsMap) Type() (Type, error) {
	if !am.columnTypeSet {
		return "", errors.New("arg type not set")
	}
	return am.columnType, nil
}

func (am *ArgsMap) Name() (string, error) {
	if !am.nameSet {
		return "", errors.New("arg name not set")
	}
	return am.name, nil
}

func (am *ArgsMap) Nullable() (bool, error) {
	if !am.nullableSet {
		return am.nullable, errors.New("arg nullable not set")
	}
	return am.nullable, nil
}

func (am *ArgsMap) Length() (uint32, error) {
	if !am.lengthSet {
		return am.length, errors.New("arg length not set")
	}
	return am.length, nil
}
