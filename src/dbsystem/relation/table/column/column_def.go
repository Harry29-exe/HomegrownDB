package column

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	. "HomegrownDB/dbsystem/relation/dbobj"
)

// Def describes column config and provides segparser and serializer
type Def interface {
	bparse.Serializable
	Name() string
	Nullable() bool
	Id() OID
	Order() Order
	CType() hgtype.TypeData

	DefaultValue() []byte
	//// Serialize should save all important Data to byte stream.
	//// It has to start with MdString of column.Tag.
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

// Order describes order of column in table
type Order = uint16

// InnerOrder describes order of column in tuple
type InnerOrder = uint16

func Deserialize(deserializer *bparse.Deserializer) WDef {
	col := &column{}
	col.Deserialize(deserializer)
	return col
}

func NewDefinition(name string, nullable bool, columnType hgtype.TypeData) WDef {
	return &column{
		name:     name,
		nullable: nullable,
		id:       0,
		order:    0,
		hgType:   columnType,
	}
}
