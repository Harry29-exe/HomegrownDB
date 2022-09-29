package column

import "HomegrownDB/dbsystem/ctype"

type Id = uint32

// Def describes column properties and provides segparser and serializer
type Def interface {
	Name() string
	Nullable() bool
	Id() Id
	Order() Order
	Type() ctype.Type
	CType() ctype.CType

	// Serialize should save all important Data to byte stream.
	// It has to start with MdString of column.Type.
	Serialize() []byte
	// Deserialize takes the same Data that Serialize returned
	// and set this column definitions to match given Data
	Deserialize(data []byte) (subsequent []byte)
}

type WDef interface {
	Def
	SetId(id Id)
	SetOrder(order Order)
}

// Order describes order of column in table
type Order = uint16

func Serialize(data []byte) (col WDef, subsequent []byte) {
	col = &column{}
	subsequent = col.Deserialize(data)
	return
}

func NewDefinition(args Args) WDef {
	//todo implement me
	panic("Not implemented")
}
