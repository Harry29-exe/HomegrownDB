package sqlparser

type TreeParentType = uint16

const (
	SELECT TreeParentType = iota
	INSERT
	UPDATE
	DELETE
)

type TreeParent interface {
	Name() TreeParentType
	Value() any
}
