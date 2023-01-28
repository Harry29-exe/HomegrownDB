package dbobj

import "math"

type OID = uint64

const (
	InvalidOID = 0
	MaxOID     = math.MaxUint64
)

type Obj interface {
	OID() OID
}
