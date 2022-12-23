package column

import "HomegrownDB/dbsystem/hgtype"

var _ WDef = &column{}

type column struct {
	name     string
	nullable bool
	id       Id
	order    Order
	hgType   hgtype.TypeData
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

func (c *column) CType() hgtype.TypeData {
	return c.hgType
}

func (c *column) DefaultValue() []byte {
	return nil
}

func (c *column) Serialize() []byte {
	//TODO implement me
	panic("implement me")
}

func (c *column) Deserialize(data []byte) (subsequent []byte) {
	//TODO implement me
	panic("implement me")
}
