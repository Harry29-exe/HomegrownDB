package column

type Type = string

// Definition describes column properties and provides parser
type Definition interface {
	Name() string
	Nullable() bool

	DataParser() DataParser
	DataSerializer() DataSerializer

	// Serialize should save all important data to byte stream.
	// It has to start with MdString of column.Type.
	Serialize() []byte
	// Deserialize takes the same data that Serialize returned
	// and set this column definitions to match given data
	Deserialize(data []byte)
}

type DataParser interface {
	Skip(data []byte) []byte
	Parse(data []byte) (Value, []byte)
}

type DataSerializer interface {
	Serialize(data []byte) ([]byte, error)
	SerializeValue(value *any) ([]byte, error)
	//todo method serialize as lob while standard Serialize return error if value is to large to be serialized in standard way
}

type Value interface {
	AsBytes() []byte
	Value() any
	IsNull() bool

	EqualsBytes(value []byte) bool
	Equals(value *any) bool

	// CompareBytes returns 0 if this.value == value,
	// -1 if this.value < value and 1 if this.value > value
	CompareBytes(value []byte) int
	// Compare returns 0 if this.value == value,
	// -1 if this.value < value and 1 if this.value > value
	Compare(value *any) int
}
