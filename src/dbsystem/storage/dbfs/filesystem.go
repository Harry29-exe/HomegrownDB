package dbfs

import (
	"HomegrownDB/dbsystem/dbobj"
	"errors"
	"fmt"
	"os"
)

type FS interface {
	PageObjectFS
	PropertiesFS
	DBInitializerFS
	Truncate(path string, newSize int64) error
	Open(path string) (FileLike, error)
}

type PropertiesFS interface {
	ReadConfigFile() ([]byte, error)
	SaveConfigFile(config []byte) error
	ReadPropertiesFile() ([]byte, error)
	SavePropertiesFile(properties []byte) error
}

type DBInitializerFS interface {
	InitDBSystemDirs() error
	InitDBSystemConfigAndProps(configData []byte, propertiesData []byte) error

	DestroyDB() error
}

type PageObjectFS interface {
	OpenPageObjectFile(oid dbobj.OID) (FileLike, error)
	OpenPageObjectDef(oid dbobj.OID) (FileLike, error)
	// InitNewRelationDir create directory and all files(empty) for given reldef id
	InitNewPageObjectDir(oid dbobj.OID) error
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

func (fs *StdFS) Truncate(path string, newSize int64) error {
	return os.Truncate(path, newSize)
}

func (fs *StdFS) Open(path string) (FileLike, error) {
	return os.OpenFile(path, os.O_RDWR, os.ModeType)
}

// -------------------------
//      FS
// -------------------------

func (fs *StdFS) ReadConfigFile() ([]byte, error) {
	path := fs.RootPath + "/" + ConfigFilename
	return fs.readFile(path)
}

// -------------------------
//      PropertiesFS
// -------------------------

func (fs *StdFS) SaveConfigFile(config []byte) error {
	configPath := Path(fs.RootPath, ConfigFilename)
	return fs.overrideFileData(configPath, config)
}

func (fs *StdFS) ReadPropertiesFile() ([]byte, error) {
	path := Path(fs.RootPath, PropertiesFilename)
	return fs.readFile(path)
}

func (fs *StdFS) SavePropertiesFile(propertiesData []byte) error {
	propsPath := Path(fs.RootPath, PropertiesFilename)
	return fs.overrideFileData(propsPath, propertiesData)
}

func (fs *StdFS) InitDBSystemDirs() error {
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

func (fs *StdFS) InitDBSystemConfigAndProps(configData []byte, propertiesData []byte) error {
	err := fs.createFile(Path(fs.RootPath, ConfigFilename), configData)
	if err != nil {
		return err
	}
	err = fs.createFile(Path(fs.RootPath, PropertiesFilename), propertiesData)
	if err != nil {
		return err
	}
	return nil
}

// -------------------------
//      DBInitializerFS
// -------------------------

func (fs *StdFS) DestroyDB() error {
	return os.RemoveAll(fs.RootPath)
}

// -------------------------
//      PageObjectFS
// -------------------------

func (fs *StdFS) OpenPageObjectFile(oid dbobj.OID) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.RootPath, RelationsDirname, oid, DataFilename)
	file, err := os.OpenFile(filepath, os.O_RDWR, os.ModeType)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) OpenPageObjectDef(oid dbobj.OID) (FileLike, error) {
	filepath := fmt.Sprintf("%s/%s/%d/%s", fs.RootPath, RelationsDirname, oid, DefinitionFilename)
	file, err := os.OpenFile(filepath, os.O_RDWR, os.ModeType)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs *StdFS) InitNewPageObjectDir(oid dbobj.OID) error {
	path := fmt.Sprintf("%s/%s/%d", fs.RootPath, RelationsDirname, oid)
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

func (fs *StdFS) overrideFileData(configPath string, config []byte) error {
	file, err := fs.Open(configPath)
	if err != nil {
		return err
	}
	if stat, err := file.Stat(); err != nil {
		return err
	} else if stat.Size() > int64(len(config)) {
		err = fs.Truncate(configPath, int64(len(config)))
		if err != nil {
			return err
		}
	}

	if _, err = file.Write(config); err != nil {
		return err
	}
	return nil
}

func (fs *StdFS) readFile(path string) (data []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err2 := file.Close()
		if err == nil {
			err = err2
		}
	}()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data = make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (fs *StdFS) createFile(path string, data []byte) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if err2 := file.Close(); err == nil {
			err = err2
		}
	}()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
