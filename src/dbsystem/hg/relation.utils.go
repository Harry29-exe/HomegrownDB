package hg

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"fmt"
)

func SerializeRel(rel relation.Relation) []byte {
	s := bparse.NewSerializer()
	switch rel.Kind() {
	case relation.TypeTable:
		rel.(table.Definition).Serialize(s)
	case relation.TypeFsm:
		fsm.SerializeFSM(rel.(*fsm.FreeSpaceMap), s)
	case relation.TypeVm:
		//todo implement me
		panic("Not implemented")
	case relation.TypeIndex:
		//todo implement me
		panic("Not implemented")
	default:
		panic(fmt.Sprintf("illegal state (unknown relation type: %+v)", rel))
	}

	return s.GetBytes()
}

func openFileAndRegisterRelationIO(relation relation.Relation, ctx *loaderCtx) error {
	file, err := ctx.FS.OpenRelationDataFile(relation)
	if err != nil {
		return err
	}
	pageIO, err := pageio.NewPageIO(file)
	if err != nil {
		return err
	}
	ctx.PageIOStore.Register(relation.RelationID(), pageIO)
	return nil
}

func readRelationFile(id relation.ID, ctx *loaderCtx) ([]byte, error) {
	file, err := ctx.FS.OpenRelationDef(id)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buff := make([]byte, stat.Size())
	_, err = file.Read(buff)
	return buff, err
}
