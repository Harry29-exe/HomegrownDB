package dbfs

import (
	"HomegrownDB/dbsystem/relation"
	"errors"
	"fmt"
	"os"
)

type FS interface {
	RelationFS
	PropertiesFS
	DBInitializerFS
}

type PropertiesFS interface {
	ReadConfigFile() ([]byte, error)
	ReadPropertiesFile() ([]byte, error)
}

type DBInitializerFS interface {
	InitDBSystem() error
}

type RelationFS interface {
	OpenRelationDataFile(relation relation.Relation) (FileLike, error)
	OpenRelationDef(relation relation.ID) (FileLike, error)
	CreateRelationDir(relation relation.Relation) error
}

func LoadFS(rootPath string) (FS, error) {
	fileStat, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	} else if !fileStat.IsDir() {
		return nil, errors.New("root path is not a directory")
	}
	return &StdFS{RootPath: rootPath}, nil
}

func CreateFS(rootPath string) (FS, error) {
	err := os.Mkdir(rootPath, permission)
	if err != nil {
		return nil, err
	}
	return &StdFS{RootPath: rootPath}, nil
}

type StdFS struct {
	RootPath string
}

const permission = 0777 //todo check this permissions

// -------------------------
//      PropertiesFS
// -------------------------

func (fs *StdFS) ReadConfigFile() ([]byte, error) {
	path := fs.RootPath + "/" + ConfigFilename
	return fs.readFile(path)
}

func (fs *StdFS) ReadPropertiesFile() ([]byte, error) {
	path := fs.RootPath + "/" + PropertiesFilename
	return fs.readFile(path)
}

// -------------------------
//      DBInitializerFS
// -------------------------

func (fs *StdFS) InitDBSystem() error {
	dirPaths := []string{
		fs.RootPath + "/" + RelationsDirname,
	}
	for _, path := range dirPaths {
		err := os.Mkdir(path, permission)
		if err != nil {
			return err
		}
	}
	return nil
}

// -------------------------
//      RelationFS
// -------------------------

func (fs *StdFS) OpenRelationDataFile(relation relation.Relation) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.RootPath, RelationsDirname, relation.RelationID(), DataFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) OpenRelationDef(relationId relation.ID) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.RootPath, RelationsDirname, relationId, DefinitionFilename)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) CreateRelationDir(relation relation.Relation) error {
	path := fmt.Sprintf("%s/%s/%d", fs.RootPath, RelationsDirname, relation.RelationID())
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return os.ErrExist
	}
	err = os.Mkdir(path, permission)
	if err != nil {
		return err
	}

	for _, filename := range relationDirFiles {
		file, err := os.Create(fmt.Sprintf("%s/%s", path, filename))
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// -------------------------
//      internal
// -------------------------

func (fs *StdFS) readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	return data, err
}
