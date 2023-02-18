package column

import (
	. "HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
)

// Def describes column config and provides parse and serializer
type Def interface {
	Name() string
	Nullable() bool
	Id() OID
	Order() Order
	CType() hgtype.ColType

	DefaultValue() []byte
	//// Serialize should save all important Data to byte stream.
	//// It has to start with MdString of column.ColTag.
	//Serialize() []byte
	//// Deserialize takes the same Data that Serialize returned
	//// and set this column definitions to match given Data
	//Deserialize(data []byte) (subsequent []byte)
}

type WDef interface {
	Def
	SetId(id OID)
	SetOrder(order Order)
}

type Id = uint32

// Order describes order of column in tabdef
type Order = uint16

// InnerOrder describes order of column in tuple
type InnerOrder = uint16

func NewDefinition(name string, oid OID, order Order, columnType hgtype.ColType) WDef {
	return &column{
		name:   name,
		id:     oid,
		order:  order,
		hgType: columnType,
	}
}
