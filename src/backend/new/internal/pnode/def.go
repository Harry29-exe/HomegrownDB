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

	Equal(node Node) bool
}

type node struct {
	tag        Tag
	startToken uint
	endToken   uint
}

func (p *node) Tag() Tag {
	if p == nil {
		return TagNil
	}
	return p.tag
}

func (p *node) SetTag(tag Tag) {
	p.tag = tag
}

func (p *node) StartToken() uint {
	return p.startToken
}

func (p *node) SetStartToken(index uint) {
	p.startToken = index
}

func (p *node) EndToken() uint {
	return p.endToken
}

func (p *node) SetEndToken(index uint) {
	p.endToken = index
}

type Tag = uint16

const (
	TagNil Tag = iota
	TagRawStmt
	TagSelectStmt
	TagInsertStmt
	TagAExpr
	TagResultTarget
	TagColumnRef
	TagRangeVar
	TagAStar
	TagAConst
)
