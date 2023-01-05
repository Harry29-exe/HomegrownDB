package initializer

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
)

type InitCtx struct {
	LoaderCtx
	DefaultRelations []relation.Relation
}

func InitializeDB(rootPath string, ctx *InitCtx) (dbsystem.DBSystem, error) {

}

func initDB(rootPath string, ctx *InitCtx) (dbsystem.DBSystem, error) {
	var err error
	if ctx.RootPath != "" {
		ctx.RootPath = rootPath
	}

	if ctx.FS != nil {
		ctx.FS, err = dbfs.CreateFS(ctx.RootPath)
		if err != nil {
			return nil, err
		}
	}

	if ctx.PageIOStore {

	}
}

func initFS(ctx *InitCtx) {
	fs, err := dbfs.CreateFS(ctx.RootPath)
	if err != nil {
		panic(err.Error())
	}

}

func initPageIO(ctx *InitCtx) {
	ctx.PageIOStore = pageio.NewStore(ctx.FS)
}

type initFn = func(ctx *InitCtx) error

func init() {

}

func handleErr[T any](v T, err error) T {
	return v
}
