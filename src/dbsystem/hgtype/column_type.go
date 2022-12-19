package hgtype

import "HomegrownDB/dbsystem/hgtype/toast"

type HGType = *hgType

// todo change to cType and create type HGType = *cType
type hgType struct {
	Type Type
	Reader
	Writer
	Operations
	Debug

	VarLen      bool
	ToastStatus toast.Status
}

func newHGType(valType Type, reader Reader, writer Writer, op Operations, deb Debug, varLen bool, status toast.Status) HGType {
	return &hgType{
		Type:        valType,
		Reader:      reader,
		Writer:      writer,
		Operations:  op,
		Debug:       deb,
		VarLen:      varLen,
		ToastStatus: status,
	}
}
