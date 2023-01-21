package hg

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"fmt"
)

func SerializeRel(rel relation.Relation) []byte {
	s := bparse.NewSerializer()
	switch rel.Kind() {
	case relation.TypeTable:
		rel.(table.Definition).Serialize(s)
	case relation.TypeIndex:
		//todo implement me
		panic("Not implemented")
	default:
		panic(fmt.Sprintf("illegal state (unknown relation type: %+v)", rel))
	}

	return s.GetBytes()
}
