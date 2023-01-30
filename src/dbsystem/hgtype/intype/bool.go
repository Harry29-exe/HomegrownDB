package intype

import "HomegrownDB/dbsystem/hgtype/rawtype"

func ConvBool(val bool) []byte {
	//todo implement me
	panic("Not implemented")
}

func ConvBoolValue(val bool) rawtype.Value {
	var byteBool byte = 0
	if val {
		byteBool = 1
	}
	return rawtype.Value{
		TypeTag:   rawtype.TypeBool,
		NormValue: []byte{byteBool},
	}
}
