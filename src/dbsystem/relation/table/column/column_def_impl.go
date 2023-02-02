package column

import (
	"HomegrownDB/dbsystem/hgtype"
	. "HomegrownDB/dbsystem/relation/dbobj"
)

var _ WDef = &column{}

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
