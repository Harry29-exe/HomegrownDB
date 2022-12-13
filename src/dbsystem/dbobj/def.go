package dbobj

type OID = uint64

type DBObj interface {
	OID() OID
}
