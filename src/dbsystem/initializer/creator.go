package initializer

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
	"os"
)

type CreatorProps struct {
	RootPath string // RootPath path where db will be initialized (nullable)
	RelationsToCreate []relation.Relation // Relations that DBCreator will create
}

func (c *CreatorProps) initEmptyWithDefault() error {
	if c.RootPath == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		c.RootPath = homedir + "/.HomegrownDB"
	}
	if c.RelationsToCreate == nil {

	}

	return nil
}

func InitializeDB(props CreatorProps) (dbsystem.DBSystem, error) {
	
}

func initializeRootPath(rootPath string) error {

}

func initDB(rootPath string, ctx *CreatorCtx) (dbsystem.DBSystem, error) {
	var err error
	if ctx.RootPath != "" {
		ctx.RootPath = rootPath
	} else {
		return dbsystem.
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

func initFS(ctx *CreatorCtx) {
	fs, err := dbfs.CreateFS(ctx.RootPath)
	if err != nil {
		panic(err.Error())
	}

}

func initPageIO(ctx *CreatorCtx) {
	ctx.PageIOStore = pageio.NewStore(ctx.FS)
}

type initFn = func(ctx *CreatorCtx) error

func init() {

}

func handleErr[T any](v T, err error) T {
	return v
}
