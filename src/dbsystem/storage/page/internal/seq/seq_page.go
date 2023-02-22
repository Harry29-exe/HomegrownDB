package seq

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page/internal"
)

func AsPage(seqOID reldef.OID, data []byte) Page {
	return Page{
		seqOID: seqOID,
		Bytes:  data,
	}
}

// Page is page in sys_sequences table
type Page struct {
	seqOID reldef.OID
	Bytes  []byte
}

func (p Page) PageTag() internal.PageTag {
	return internal.PageTag{
		PageId:  0, // only one page is created for each sequence
		OwnerID: p.seqOID,
	}
}
