package column

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
