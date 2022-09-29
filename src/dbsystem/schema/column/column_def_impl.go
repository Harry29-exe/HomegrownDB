package column

import "HomegrownDB/dbsystem/ctype"

var _ WDef = &column{}

type column struct {
	name     string
	nullable bool
	id       Id
	order    Order
	typeCode ctype.Type
	ctype    ctype.CType
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Nullable() bool {
	return c.nullable
}

func (c *column) Id() Id {
	return c.id
}

func (c *column) SetId(id Id) {
	c.id = id
}

func (c *column) Order() Order {
	return c.order
}

func (c *column) SetOrder(order Order) {
	c.order = order
}

func (c *column) Type() ctype.Type {
	return c.typeCode
}

func (c *column) CType() ctype.CType {
	return c.ctype
}

func (c *column) Serialize() []byte {
	//TODO implement me
	panic("implement me")
}

func (c *column) Deserialize(data []byte) (subsequent []byte) {
	//TODO implement me
	panic("implement me")
}
