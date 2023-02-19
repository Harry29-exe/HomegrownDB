package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/access/table"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
	"log"
)

type Manager interface {
	Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error)
	Delete(relation reldef.Relation) error
	AccessMngr
}

type AccessMngr interface {
	// FindByName finds relation with provided name and returns its oid,
	// if relation does not exist returns dbobj.InvalidId
	FindByName(name string) reldef.OID
	Access(oid reldef.OID, mode LockMode) reldef.Relation
	Free(relationOID reldef.OID, mode LockMode)
}

type LockMode uint8

const (
	LockNone LockMode = iota
	LockRead
	LockWrite
)

type OIDSequence interface {
	Next() reldef.OID
}

type stdManager struct {
	Buffer      buffer.SharedBuffer
	FS          dbfs.FS
	OIDSequence OIDSequence

	nameMap map[string]reldef.OID
	cache   cache
}

var _ Manager = &stdManager{}

func (s *stdManager) Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error) {
	if relation.OID() == hglib.InvalidOID {
		err := s.initRelation(relation)
		if err != nil {
			return relation, err
		}
	}
	if err := s.createRelationOnDisc(relation); err != nil {
		return nil, err
	}

	s.cache.relations[relation.OID()] = relation
	s.nameMap[relation.Name()] = relation.OID()

	switch relation.Kind() {
	case reldef.TypeTable:
		return relation, s.createTableInSysTables(relation.(tabdef.Definition), tx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (s *stdManager) initRelation(relation reldef.Relation) error {
	relation.InitRel(s.OIDSequence.Next(), s.OIDSequence.Next(), s.OIDSequence.Next())
	switch relation.Kind() {
	case reldef.TypeTable:
		tableDef := relation.(tabdef.Definition)
		for _, col := range tableDef.Columns() {
			(col.(tabdef.ColumnDefinition)).SetId(s.OIDSequence.Next())
		}
	}

	return nil
}

func (s *stdManager) createTableInSysTables(definition tabdef.Definition, tx tx.Tx) error {
	tuple, err := systable.RelationsOps.TableAsRelationsRow(definition, tx)
	if err != nil {
		return err
	}
	err = table.Insert(tuple, tx, systable.RelationsTableDef(), fsm.NewFSM(systable.HGRelationsFsmOID, s.Buffer), s.Buffer)
	if err != nil {
		return err
	}
	columns := systable.ColumnsOps.DataToRows(definition.OID(), definition.Columns(), tx)
	for _, columnRow := range columns {
		err = table.Insert(columnRow, tx, systable.ColumnsTableDef(), fsm.NewFSM(systable.HGColumnsFsmOID, s.Buffer), s.Buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *stdManager) createRelationOnDisc(relation reldef.Relation) error {
	if err := s.FS.InitNewPageObjectDir(relation.OID()); err != nil {
		return err
	}
	if err := s.FS.InitNewPageObjectDir(relation.FsmOID()); err != nil {
		return err
	}
	if err := s.FS.InitNewPageObjectDir(relation.VmOID()); err != nil {
		return err
	}

	if err := fsm.InitFreeSpaceMapFile(relation.FsmOID(), s.FS); err != nil {
		return err
	}

	//todo add err handling (delete all created on fail)
	return nil
}

func (s *stdManager) Delete(relation reldef.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) FindByName(name string) reldef.OID {
	oid, ok := s.nameMap[name]
	if !ok {
		return reldef.InvalidRelId
	}
	return oid
}

func (s *stdManager) Access(oid reldef.OID, mode LockMode) reldef.Relation {
	log.Printf("waring: locking relations is not done")
	return s.cache.relations[oid] //todo locking
}

func (s *stdManager) Free(relationOID reldef.OID, mode LockMode) {
	//todo locking
	log.Printf("waring: locking relations is not done")
}
