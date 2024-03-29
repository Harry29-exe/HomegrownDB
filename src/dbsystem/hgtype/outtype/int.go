package outtype

import (
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/lib/bparse"
	"log"
)

func AsInt(val []byte, t rawtype.Tag) int64 {
	switch t {
	case rawtype.TypeInt8:
		return bparse.Parse.Int8(val)
	default:
		log.Panicf("not supported type %s", t.ToStr())
		return 0
	}
}
