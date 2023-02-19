package seq

import "HomegrownDB/dbsystem/reldef"

// Page is page in sys_sequences table
type Page struct {
	seqOID reldef.OID
	bytes  []byte
}

type Tuple struct {
}
