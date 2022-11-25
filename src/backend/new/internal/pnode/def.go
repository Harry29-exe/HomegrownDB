package pnode

type Node interface {
	Tag() Tag
	SetTag(tag Tag)
	// StartToken index of first token of this node
	StartToken() uint
	SetStartToken(index uint)
	// EndToken index of first token after this node
	EndToken() uint
	SetEndToken(index uint)
}

var _ Node = &node{}

type node struct {
	tag        Tag
	startToken uint
	endToken   uint
}

func (p *node) Tag() Tag {
	//TODO implement me
	panic("implement me")
}

func (p *node) SetTag(tag Tag) {
	//TODO implement me
	panic("implement me")
}

func (p *node) StartToken() uint {
	//TODO implement me
	panic("implement me")
}

func (p *node) SetStartToken(index uint) {
	//TODO implement me
	panic("implement me")
}

func (p *node) EndToken() uint {
	//TODO implement me
	panic("implement me")
}

func (p *node) SetEndToken(index uint) {
	//TODO implement me
	panic("implement me")
}

type Tag = uint16

const (
	TagSelectStmt Tag = iota
	TagInsertStmt
	TagExpr
	TagResultTarget
	TagColumnRef
	TagRangeVar
)
