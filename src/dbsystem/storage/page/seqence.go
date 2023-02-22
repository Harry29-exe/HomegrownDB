package page

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page/internal/seq"
)

type SequencePage = seq.Page

func AsSequencePage(seqOID reldef.OID, data []byte) SequencePage {
	return seq.AsPage(seqOID, data)
}
