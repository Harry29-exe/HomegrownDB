package column

import "HomegrownDB/dbsystem/hgtype"

// Def describes column config and provides segparser and serializer
type Def interface {
	Name() string
	Nullable() bool
	Id() Id
	Order() Order
	CType() hgtype.TypeData

	// Serialize should save all important Data to byte stream.
	// It has to start with MdString of column.Tag.
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

type Id = uint32

// Order describes order of column in table
type Order = uint16

// InnerOrder describes order of column in tuple
type InnerOrder = uint16

func Serialize(data []byte) (col WDef, subsequent []byte) {
	col = &column{}
	subsequent = col.Deserialize(data)
	return
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
