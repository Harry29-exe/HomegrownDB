package factory

type Args interface {
	GetArg(code ArgCode) any
}

type ArgCode uint16

type ArgsMap struct {
	args map[ArgCode]any
}

func (am *ArgsMap) GetArg(code ArgCode) any {
	return am.args[code]
}

func (am *ArgsMap) AddArgs(code ArgCode, value any) {
	am.args[code] = value
}
