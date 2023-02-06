package testinfr

import (
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"sync"
)

func NewRelationAccessManager(relations ...reldef.Relation) relation.AccessMngr {
	mngr := new(RelationsAccessManager)
	mngr.Relations = map[reldef.OID]reldef.Relation{}
	mngr.LockMap = map[reldef.OID]*sync.RWMutex{}
	for _, rel := range relations {
		mngr.Relations[rel.OID()] = rel
		mngr.LockMap[rel.OID()] = &sync.RWMutex{}
	}
	return mngr
}

type RelationsAccessManager struct {
	Relations map[reldef.OID]reldef.Relation
	LockMap   map[reldef.OID]*sync.RWMutex
}

func (r *RelationsAccessManager) FindByName(name string) reldef.OID {
	for _, rel := range r.Relations {
		if rel.Kind() == reldef.TypeTable &&
			rel.(tabdef.RDefinition).Name() == name {
			return rel.OID()
		}
	}
	return reldef.InvalidRelId
}

func (r *RelationsAccessManager) Access(oid reldef.OID, mode relation.LockMode) reldef.Relation {
	switch mode {
	case relation.LockRead:
		r.LockMap[oid].RLock()
	case relation.LockWrite:
		r.LockMap[oid].Lock()
	case relation.LockNone:
		break
	}
	return r.Relations[oid]
}

func (r *RelationsAccessManager) Free(relationOID reldef.OID, mode relation.LockMode) {
	switch mode {
	case relation.LockRead:
		r.LockMap[relationOID].RUnlock()
	case relation.LockWrite:
		r.LockMap[relationOID].Unlock()
	case relation.LockNone:
		break
	}
}

var _ relation.AccessMngr = &RelationsAccessManager{}
