package column

import "HomegrownDB/dbsystem/ctype"

var _ WDef = &column{}

type column struct {
	name     string
	nullable bool
	id       Id
	typeCode ctype.Type
	ctype    ctype.CType
}

func (c *column) SetColumnId(id Id) {
	//TODO implement me
	panic("implement me")
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Nullable() bool {
	return c.nullable
}

func (c *column) GetColumnId() Id {
	return c.id
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
