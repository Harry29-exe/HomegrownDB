package reldef

import "HomegrownDB/dbsystem/hgtype/rawtype"

type SequenceDef interface {
	RSequenceDef
	SetTypeTag(typeTag rawtype.Tag)
	SetSeqStart(seqStart int64)
	SetSeqIncrement(seqIncrement int64)
	SetSeqMax(seqMax int64)
	SetSeqMin(seqMin int64)
	SetSeqCacheSize(seqCacheSize int64)
	SetSeqCycle(seqCycle bool)
}

type RSequenceDef interface {
	Relation
	TypeTag() rawtype.Tag
	SeqStart() int64
	SeqIncrement() int64
	SeqMax() int64
	SeqMin() int64
	SeqCacheSize() int64
	SeqCycle() bool
}

func NewSequenceDef(
	relation Relation,
	typeTag rawtype.Tag,
	seqStart int64,
	seqIncrement int64,
	seqMax int64,
	seqMin int64,
	seqCacheSize int64,
	seqCycle bool,
) SequenceDef {
	return &Sequence{
		Relation:     relation,
		typeTag:      typeTag,
		seqStart:     seqStart,
		seqIncrement: seqIncrement,
		seqMax:       seqMax,
		seqMin:       seqMin,
		seqCacheSize: seqCacheSize,
		seqCycle:     seqCycle,
	}
}

type Sequence struct {
	Relation
	typeTag      rawtype.Tag
	seqStart     int64
	seqIncrement int64
	seqMax       int64
	seqMin       int64
	seqCacheSize int64
	seqCycle     bool
}

func (s *Sequence) TypeTag() rawtype.Tag {
	return s.typeTag
}

func (s *Sequence) SeqStart() int64 {
	return s.seqStart
}

func (s *Sequence) SeqIncrement() int64 {
	return s.seqIncrement
}

func (s *Sequence) SeqMax() int64 {
	return s.seqMax
}

func (s *Sequence) SeqMin() int64 {
	return s.seqMin
}

func (s *Sequence) SeqCacheSize() int64 {
	return s.seqCacheSize
}

func (s *Sequence) SeqCycle() bool {
	return s.seqCycle
}

func (s *Sequence) SetTypeTag(typeTag rawtype.Tag) {

}

func (s *Sequence) SetSeqStart(seqStart int64) {
	s.seqStart = seqStart
}

func (s *Sequence) SetSeqIncrement(seqIncrement int64) {
	s.seqIncrement = seqIncrement
}

func (s *Sequence) SetSeqMax(seqMax int64) {
	s.seqMax = seqMax
}

func (s *Sequence) SetSeqMin(seqMin int64) {
	s.seqMin = seqMin
}

func (s *Sequence) SetSeqCacheSize(seqCacheSize int64) {
	s.seqCacheSize = seqCacheSize
}

func (s *Sequence) SetSeqCycle(seqCycle bool) {
	s.seqCycle = seqCycle
}
