package dbfs

const (
	ConfigFilename     = "config.json"
	PropertiesFilename = "properties.json"

	RelationsDirname   = "relations"
	DataFilename       = "data.hdb"
	DefinitionFilename = "def.hdb"
)

var relationDirFiles = [...]string{
	DataFilename,
	DefinitionFilename,
}
