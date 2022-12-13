package ctype

// todo change to cType and create type CType = *cType
type CType struct {
	Type Type
	Reader
	Writer
	Operations
	Debug

	VarLen      bool
	ToastStatus ToastStatus
}

func newCType(valType Type, reader Reader, writer Writer, op Operations, deb Debug, varLen bool, status ToastStatus) *CType {
	return &CType{
		Type:        valType,
		Reader:      reader,
		Writer:      writer,
		Operations:  op,
		Debug:       deb,
		VarLen:      varLen,
		ToastStatus: status,
	}
}
