package ctype

type CType struct {
	Reader
	Writer
	Operations
	Debug

	VarLen      bool
	ToastStatus ToastStatus
}

func newCType(reader Reader, op Operations, deb Debug, varLen bool, status ToastStatus) *CType {
	return &CType{
		reader,
		nil,
		op,
		deb,
		varLen,
		status,
	}
}
