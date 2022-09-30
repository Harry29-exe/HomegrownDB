package exenode

import (
	"HomegrownDB/backend/qrow"
)

type Insert struct {
}

func (i *Insert) SetSource(source []ExeNode) {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Init(options InitOptions) qrow.RowBuffer {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) HasNext() bool {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Next() qrow.Row {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) NextBatch() []qrow.Row {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) All() []qrow.Row {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Free() {
	//TODO implement me
	panic("implement me")
}
