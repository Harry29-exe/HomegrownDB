package access

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hglib"
)

func InstallModule(mode hglib.Mode, deps ModuleDeps) Module {
	buffSize := deps.ConfigModule.Config().SharedBufferSize
	buff := buffer.NewSharedBuffer(buffSize, deps.StorageModule.PageIOStore())

	_ = buff
	//todo implement me
	panic("Not implemented")
}
