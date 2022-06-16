package column

type Value interface {
	AsBytes() []byte
	Value() any
	IsNull() bool

	Equals(value *any) bool
	// Compare returns 0 if this.value == value,
	// -1 if this.value < value and 1 if this.value > value
	Compare(value *any) int
}
