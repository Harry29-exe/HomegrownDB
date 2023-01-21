package dbobj

type OID = uint64

type Obj interface {
	OID() OID
}
