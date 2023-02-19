package tabdef

import (
	. "HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype"
)

// ColumnRDefinition describes column config and provides parse and serializer
type ColumnRDefinition interface {
	Name() string
	Nullable() bool
	Id() OID
	Order() Order
	CType() hgtype.ColType

	DefaultValue() []byte
}

type ColumnDefinition interface {
	ColumnRDefinition
	SetId(id OID)
	SetOrder(order Order)
}

// Order describes order of column in tabdef
type Order = uint16

// InnerOrder describes order of column in tuple
type InnerOrder = uint16

func NewColumnDefinition(name string, oid OID, order Order, columnType hgtype.ColType) ColumnDefinition {
	return &column{
		name:   name,
		id:     oid,
		order:  order,
		hgType: columnType,
	}
}

var _ ColumnDefinition = &column{}

type column struct {
	name     string
	nullable bool
	id       OID
	order    Order
	hgType   hgtype.ColType
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Nullable() bool {
	return c.nullable
}

func (c *column) Id() OID {
	return c.id
}

func (c *column) SetId(id OID) {
	c.id = id
}

func (c *column) Order() Order {
	return c.order
}

func (c *column) SetOrder(order Order) {
	c.order = order
}

func (c *column) CType() hgtype.ColType {
	return c.hgType
}

func (c *column) DefaultValue() []byte {
	return nil
}
