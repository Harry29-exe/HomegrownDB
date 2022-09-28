package ctype

func convInputStr(input string, cType CType) ([]byte, error) {
	switch cType {
	case Str:
		return []byte(input), nil
	default:
		return nil, NewInputConvErr(InputTypeStr, input, cType)
	}
}
