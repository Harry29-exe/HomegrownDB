package hglib

type Module interface {
	Shutdown() error
}

type Mode uint8

const (
	InstallerModeDB Mode = iota
	InstallerModeTest
)

type ModuleInstallerArgs struct {
	Mode       Mode
	RootPath   string // RootPath path where db will be initialized (nullable)
	BufferSize uint
}
