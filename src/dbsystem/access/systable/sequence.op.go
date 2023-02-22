package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var SequencesOps = sequencesOps{}

type sequencesOps struct{}

func (sequencesOps) DataToRow(sequence reldef.SequenceDef, tx tx.Tx) page.WTuple {
	builder := internal.NewTupleBuilder(sequenceDef)

	builder.WriteValue(intype.ConvInt8Value(int64(sequence.OID())))
	builder.WriteValue(intype.ConvInt8Value(int64(sequence.TypeTag())))
	builder.WriteValue(intype.ConvInt8Value(sequence.SeqStart()))
	builder.WriteValue(intype.ConvInt8Value(sequence.SeqIncrement()))
	builder.WriteValue(intype.ConvInt8Value(sequence.SeqMax()))
	builder.WriteValue(intype.ConvInt8Value(sequence.SeqMin()))
	builder.WriteValue(intype.ConvInt8Value(sequence.SeqCacheSize()))
	builder.WriteValue(intype.ConvBoolValue(sequence.SeqCycle()))

	return builder.Tuple(tx)
}

func (sequencesOps) RowToData(tuple page.RTuple) FutureSequence {
	oid := reldef.OID(internal.GetInt8(SequencesOrderOID, tuple))
	typeTag := rawtype.Tag(internal.GetInt8(SequencesOrderTypeTag, tuple))
	start := internal.GetInt8(SequencesOrderSeqStart, tuple)
	increment := internal.GetInt8(SequencesOrderSeqIncrement, tuple)
	max := internal.GetInt8(SequencesOrderSeqMax, tuple)
	min := internal.GetInt8(SequencesOrderSeqMin, tuple)
	cacheSize := internal.GetInt8(SequencesOrderSeqCacheSize, tuple)
	seqCycle := internal.GetBool(SequencesOrderSeqCycle, tuple)

	return FutureSequence{
		oid,
		typeTag,
		start,
		increment,
		max,
		min,
		cacheSize,
		seqCycle,
	}
}

type FutureSequence struct {
	oid          OID
	typeTag      rawtype.Tag
	seqStart     int64
	seqIncrement int64
	seqMax       int64
	seqMin       int64
	seqCacheSize int64
	seqCycle     bool
}

func (f FutureSequence) OID() reldef.OID {
	return f.oid
}

func (f FutureSequence) Init(relation reldef.Relation) reldef.SequenceDef {
	return reldef.NewSequenceDef(
		relation,
		f.typeTag,
		f.seqStart,
		f.seqIncrement,
		f.seqMax,
		f.seqMin,
		f.seqCacheSize,
		f.seqCycle,
	)
}
