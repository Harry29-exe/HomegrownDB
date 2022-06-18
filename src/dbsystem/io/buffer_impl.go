package io

type buffer struct {
	pageIdBufferId  map[PageTag]BufferId
	descriptorArray []bufferDescriptor
	pageBufferArray []byte
}

func (b *buffer) RPage(tag PageTag) (RPage, error) {
	panic("Not implemented")
	//pageDescriptor, ok := b.pageIdBufferId[tag]
	//if !ok {
	//
	//}

}

func (b *buffer) WPage(id PageId) (WPage, error) {
	panic("Not implemented")
}

func (b *buffer) ReleaseWPage(page WPage) {
	panic("Not implemented")
}

func (b *buffer) ReleaseRPage(page RPage) {
	panic("Not implemented")
}

func (b *buffer) fetchPage(tag PageTag) error {
	panic("Not implemented")
}
