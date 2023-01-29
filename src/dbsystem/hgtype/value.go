package hgtype

type Value struct {
	TypeTag   Tag    // Tag of type of Value
	NormValue []byte // Value in normalized form
}
