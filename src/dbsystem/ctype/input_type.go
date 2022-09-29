package ctype

import (
	"fmt"
)

type InputType = uint8

const (
	InputTypeInt   InputType = iota
	InputTypeFloat InputType = iota
	InputTypeStr   InputType = iota
)

func ConvInput(input any, cType Type) ([]byte, error) {
	switch input.(type) {
	case int64:
		return convInputInt(input.(int64), cType)
	case float64:
		return convInputFloat(input.(float64), cType)
	case string:
		return convInputStr(input.(string), cType)
	default:
		panic(fmt.Sprintf("not implemented unexpected type with value %+v", input))
	}
}

func ConvInputBuff(dst []byte, input any, cType Type) ([]byte, error) {
	val, err := ConvInput(input, cType)
	if err != nil {
		return nil, err
	}

	return append(dst, val...), nil
}

func InputTypeToStr(inputType InputType) string {
	return inputTypeNameMap[inputType]
}

var inputTypeNameMap = map[InputType]string{
	InputTypeInt:   "InputTypeInt",
	InputTypeFloat: "InputTypeFloat",
	InputTypeStr:   "InputTypeStr",
}
