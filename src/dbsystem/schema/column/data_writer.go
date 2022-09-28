package column

// DataWriter to serialize data incoming from either from outside
// or data processed by internal functions, it's usually obtained from
// column.Definition
type DataWriter interface {
	// SerializeValue takes value check if it's of supported type for serialization and
	// returns DataToSave, if error occurred it returns nil, error
	SerializeValue(value any) (DataToSave, error)
	// SerializeInput takes string and attempt to convert it into adequate data type,
	// so it can be serialized
	SerializeInput(value string) (DataToSave, error)

	StorageType() storageType
	Write(tupleBuffer []byte, data []byte)
	WriteBg()
}

type storageType = uint8

const (
	StorageTypeTuple storageType = iota
	StorageTypeBg
	StorageTypeLOB
)
