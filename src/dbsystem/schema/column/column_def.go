package column

import "HomegrownDB/dbsystem/ctype"

type Type = string
type Id = uint16

// Definition describes column properties and provides segparser and serializer
type Definition interface {
	Name() string
	Nullable() bool
	GetColumnId() Id
	Type() Type
	CType() ctype.Type

	DataWriter() DataWriter

	// Serialize should save all important Data to byte stream.
	// It has to start with MdString of column.Type.
	Serialize() []byte
	// Deserialize takes the same Data that Serialize returned
	// and set this column definitions to match given Data
	Deserialize(data []byte) (subsequent []byte)
}

type WDefinition interface {
	Definition
	SetColumnId(id Id)
}

// OrderId describes order of column in table
type OrderId = uint16
