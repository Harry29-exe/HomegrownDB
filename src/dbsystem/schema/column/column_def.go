package column

import "HomegrownDB/dbsystem/ctype"

// Def describes column properties and provides segparser and serializer
type Def interface {
	Name() string
	Nullable() bool
	Id() Id
	Order() Order
	InnerOrder() InnerOrder
	Type() ctype.Type
	CType() *ctype.CType

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
	SetInnerOrder(order InnerOrder)
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

func NewDefinition(name string, nullable bool, cType ctype.Type, args ctype.Args) (WDef, error) {
	c := &column{
		name:     name,
		nullable: nullable,
		id:       0,
		order:    0,
		typeCode: cType,
		ctype:    nil,
	}
	CType, err := ctype.CreateCType(cType, args)
	if err != nil {
		return nil, err
	}
	c.ctype = CType
	return c, nil
}
