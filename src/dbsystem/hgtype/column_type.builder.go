package hgtype

import "HomegrownDB/dbsystem/dberr"

type ArgumentMap = map[Argument]any

type Argument string

const (
	ArgLength     Argument = "ArgLength"
	ArgVaryingLen Argument = "ArgVaryingLen"
	ArgUTF8       Argument = "ArgUTF8"
)

func Create(t Type, args ArgumentMap) (HGType, error) {
	switch t {
	case TypeStr:
		return createStr(args)
	case TypeInt8:
		return createInt8(args)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func createStr(args ArgumentMap) (HGType, error) {
	parser := argParser{args: args}

	length := parser.Length()
	varyingLen := parser.VaryingLen()
	utf8 := parser.UTF8()

	if err := parser.Error(); err != nil {
		return nil, err
	}

	strArgs := StrArgs{
		Length:     uint32(length),
		VaryingLen: varyingLen,
		UTF8:       utf8,
	}

	return NewStr(strArgs), nil
}

func createInt8(args ArgumentMap) (HGType, error) {
	return NewInt8(Int8Args{}), nil
}

// -------------------------
//      ArgParser
// -------------------------

type argParser struct {
	args ArgumentMap
	err  error
}

func (a argParser) Error() error {
	return a.err
}

func (a argParser) Length() int {
	length, ok := a.args[ArgLength]
	if !ok {
		return 1
	}
	switch length.(type) {
	case int:
		return length.(int)
	default:
		a.err = dberr.NewInvalidColumnArg("ArgLength", "Arg Length must be of type integer")
		return -1
	}
}

func (a argParser) VaryingLen() bool {
	varyingLen, ok := a.args[ArgVaryingLen]
	if !ok {
		return true
	}
	switch varyingLen.(type) {
	case bool:
		return varyingLen.(bool)
	default:
		a.err = dberr.NewInvalidColumnArg("ArgVaryingLen", "Arg VaryingLen must be of type integer")
		return true
	}
}

func (a argParser) UTF8() bool {
	utf8, ok := a.args[ArgUTF8]
	if !ok {
		return true
	}
	switch utf8.(type) {
	case bool:
		return utf8.(bool)
	default:
		a.err = dberr.NewInvalidColumnArg("ArgUTF8", "Arg UTF8 must be of type integer")
		return true
	}
}
