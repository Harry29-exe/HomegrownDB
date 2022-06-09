package page

type buffer struct {
	pageIdBufferId  map[Tag]BufferId
	descriptorArray []bufferDescriptor
	pageBufferArray []byte
}

func (b *buffer) RPage(tag Tag) (RPage, error) {
	pageDescriptor, ok := b.pageIdBufferId[tag]
	if !ok {

	}

}

func (b *buffer) WPage(id Id) (WPage, error) {

}

func (b *buffer) ReleaseWPage(page WPage) {

}

func (b *buffer) ReleaseRPage(page RPage) {

}

func (b *buffer) fetchPage(tag Tag) error {

}
