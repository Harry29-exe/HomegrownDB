package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func NewManager(
	buffer buffer.SharedBuffer,
	fs dbfs.FS,
	oidSequence OIDSequence,
) Manager {

}
