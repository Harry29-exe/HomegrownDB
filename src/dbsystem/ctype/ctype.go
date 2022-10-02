package ctype

type CType struct {
	Reader
	Operations
	Debug

	VarLen      bool
	ToastStatus ToastStatus
}

func newCType(reader Reader, op Operations, deb Debug, varLen bool, status ToastStatus) *CType {
	return &CType{
		reader,
		op,
		deb,
		varLen,
		status,
	}
}
