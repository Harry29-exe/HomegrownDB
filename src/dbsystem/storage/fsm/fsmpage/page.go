package fsmpage

const (
	headerSize = 8
)

type Page struct {
	Bytes []byte
}

func (p Page) Header() []byte {
	return p.Bytes[:headerSize]
}

func (p Page) Data() []byte {
	return p.Bytes[headerSize:]
}
