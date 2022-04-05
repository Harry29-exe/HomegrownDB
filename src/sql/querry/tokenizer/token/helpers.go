package token

func IsBreak(code Code) bool {
	return breakCodes[code]
}
