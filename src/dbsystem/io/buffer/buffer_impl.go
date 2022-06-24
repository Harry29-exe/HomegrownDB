package buffer

import "HomegrownDB/dbsystem/bstructs"

type buffer struct {
	pageIdBufferId  map[bstructs.PageTag]ArrayIndex
	descriptorArray []bufferDescriptor
	pageBufferArray []byte
}

func (b *buffer) RPage(tag bstructs.PageTag) (bstructs.RPage, error) {
	panic("Not implemented")
	//pageDescriptor, ok := b.pageIdBufferId[tag]
	//if !ok {
	//
	//}

}

func (b *buffer) WPage(id bstructs.PageId) (bstructs.WPage, error) {
	panic("Not implemented")
}

func (b *buffer) ReleaseWPage(page bstructs.WPage) {
	panic("Not implemented")
}

func (b *buffer) ReleaseRPage(page bstructs.RPage) {
	panic("Not implemented")
}

func (b *buffer) fetchPage(tag bstructs.PageTag) error {
	panic("Not implemented")
}
