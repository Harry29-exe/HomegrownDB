package initializer

import (
	"HomegrownDB/dbsystem/relation"
	"os"
)

type CreatorMode uint8

const (
	DBInstaller CreatorMode = iota
	Test
)

type CreatorProps struct {
	Mode              CreatorMode
	RootPath          string              // RootPath path where db will be initialized (nullable)
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

func InitializeDB(props CreatorProps) error {
	ctx := CreatorCtx{Props: props}
	ctx.initRootPath().
		initConfigurationAndProperties().
		initDBFilesystem().
		createRelations()
	return ctx.err
}
