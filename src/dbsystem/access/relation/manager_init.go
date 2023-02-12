package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func NewManager(
	buffer buffer.SharedBuffer,
	fs dbfs.FS,
	oidSequence OIDSequence,
) (Manager, error) {
	manager := &stdManager{
		Buffer:      buffer,
		FS:          fs,
		OIDSequence: oidSequence,
		nameMap:     map[string]reldef.OID{},
		cache:       cache{},
	}
	err := manager.cache.Reload(buffer)
	if err != nil {
		return nil, err
	}
	for _, relation := range manager.cache.relations {
		manager.nameMap[relation.Name()] = relation.OID()
	}

	return manager, nil
}
