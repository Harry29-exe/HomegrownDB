package column

type Type = string

// Definition describes column properties and provides parser and serializer
type Definition interface {
	Name() string
	Nullable() bool

	DataParser() DataParser
	DataSerializer() DataSerializer

	// Serialize should save all important Data to byte stream.
	// It has to start with MdString of column.Type.
	Serialize() []byte
	// Deserialize takes the same Data that Serialize returned
	// and set this column definitions to match given Data
	Deserialize(data []byte) (subsequent []byte)
}

// OrderId describes order of column in table
type OrderId = uint16

// DataParser to parse raw data obtained from disc,
// it's usually obtained from column.Definition
type DataParser interface {
	Skip(data []byte) []byte
	Parse(data []byte) (Value, []byte)
}

// DataSerializer to serialize data incoming from either from outside
// or data processed by internal functions, it's usually obtained from
// column.Definition
type DataSerializer interface {
	// Serialize takes Data for serialization and returns DataToSave,
	// if error occurred it returns nil, error
	Serialize(data []byte) (DataToSave, error)
	// SerializeValue takes value check if it's of supported type for serialization and
	// returns DataToSave, if error occurred it returns nil, error
	SerializeValue(value any) (DataToSave, error)
}

// DataToSave data created by DataSerializer, contains data
// that should be saved to disc and information where exactly
// on disc data should be saved
type DataToSave interface {
	Data() []byte
	StorePlace() DataStoragePlace
}

func NewDataToSave(data []byte, storePlace DataStoragePlace) DataToSave {
	return &dataToSave{
		data:       data,
		storePlace: storePlace,
	}
}

type dataToSave struct {
	data       []byte
	storePlace DataStoragePlace
}

func (d *dataToSave) Data() []byte {
	return d.data
}

func (d *dataToSave) StorePlace() DataStoragePlace {
	return d.storePlace
}

type DataStoragePlace = uint8

const (
	StoreInTuple DataStoragePlace = iota
	StoreInBackground
	StoreInLob
)
