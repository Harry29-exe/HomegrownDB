package dbfs

import (
	"HomegrownDB/dbsystem/schema/relation"
	"fmt"
	"os"
)

type FS interface {
	OpenRelationDataFile(relation relation.Relation) (FileLike, error)
	OpenRelationDef(relation relation.Relation) (FileLike, error)
}

type StdFS struct {
	Root string
}

func (fs *StdFS) OpenRelationDataFile(relation relation.Relation) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.Root, RelationsDirname, relation.RelationID(), DataFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) OpenRelationDef(relation relation.Relation) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.Root, RelationsDirname, relation.RelationID(), DefinitionFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
