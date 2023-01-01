package dbfs

import (
	"HomegrownDB/dbsystem/relation"
	"fmt"
	"os"
)

type FS interface {
	OpenRelationDataFile(relation relation.Relation) (FileLike, error)
	OpenRelationDef(relation relation.ID) (FileLike, error)
}

func NewFS(rootpath string) FS {
	return &StdFS{Rootpath: rootpath}
}

type StdFS struct {
	Rootpath string
}

func (fs *StdFS) OpenRelationDataFile(relation relation.Relation) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.Rootpath, RelationsDirname, relation.RelationID(), DataFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) OpenRelationDef(relationId relation.ID) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.Rootpath, RelationsDirname, relationId, DefinitionFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
