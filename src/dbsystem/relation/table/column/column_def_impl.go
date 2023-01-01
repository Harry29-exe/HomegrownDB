package column

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	. "HomegrownDB/dbsystem/relation/dbobj"
)

var _ WDef = &column{}

type column struct {
	name     string
	nullable bool
	id       OID
	order    Order
	hgType   hgtype.TypeData
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

func (c *column) CType() hgtype.TypeData {
	return c.hgType
}

func (c *column) DefaultValue() []byte {
	return nil
}

func (c *column) Serialize(serializer *bparse.Serializer) {
	serializer.MdString(c.name)
	serializer.Bool(c.nullable)
	serializer.Uint32(uint32(c.id))
	serializer.Uint16(c.order)
	hgtype.SerializeTypeData(c.hgType, serializer)
}

func (c *column) Deserialize(deserializer *bparse.Deserializer) {
	c.name = deserializer.MdString()
	c.nullable = deserializer.Bool()
	c.id = OID(deserializer.Uint32())
	c.order = deserializer.Uint16()
	c.hgType = hgtype.DeserializeTypeData(deserializer)
}
