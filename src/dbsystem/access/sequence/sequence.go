package sequence

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/lib/bparse"
	"log"
	"math"
)

const (
	MaxINT8Increment = math.MaxInt64 / 4
)

func NextValue(sequence reldef.RSequenceDef, buffer buffer.SharedBuffer) (rawtype.Value, error) {
	seqPage, err := buffer.SeqPage(sequence.OID())
	if err != nil {
		return rawtype.Value{}, err
	}
	defer buffer.WPageRelease(seqPage.PageTag())

	switch sequence.TypeTag() {
	case rawtype.TypeInt8:
		return nextValueINT8(sequence, seqPage)
	default:
		log.Panicf("not supported type: %s", sequence.TypeTag().ToStr())
		panic("")
	}
}

func nextValueINT8(sequence reldef.RSequenceDef, seqPage page.SequencePage) (rawtype.Value, error) {
	currentVal := bparse.Parse.Int8(seqPage.Bytes)
	var value = rawtype.Value{TypeTag: rawtype.TypeInt8}
	var nextVal int64

	if sequence.SeqCycle() {
		//todo implement me (this needs to handle int wrap around)
		panic("Not implemented")
	} else {
		increment := sequence.SeqIncrement()
		nextVal = currentVal + increment
		if increment > 0 && nextVal < currentVal {
			return value, hglib.IntegerOverflowErr
		} else if increment < 0 && nextVal > currentVal {
			return value, hglib.IntegerOverflowErr
		}
		value.NormValue = bparse.Serialize.Int8(nextVal)
		copy(seqPage.Bytes, value.NormValue)
		return value, nil
	}
}
