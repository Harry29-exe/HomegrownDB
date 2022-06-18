package types

import "HomegrownDB/dbsystem/schema/column"

type prototype struct {
	columnId column.Id
}

func (c *prototype) GetColumnId() column.Id {
	return c.columnId
}

func (c *prototype) SetColumnId(id column.Id) {
	c.columnId = id
}
