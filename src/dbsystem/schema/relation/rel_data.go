package relation

type Data = *data

func NewData(dirFilePath string, dataSize int64) Data {
	return &data{
		dataSize:    dataSize,
		dirFilepath: dirFilePath,
	}
}

type data struct {
	// dataSize size of all pages in relation
	dataSize int64
	// dirFilepath path to directory in which relation
	// files exist
	dirFilepath string
}

func (d Data) IncrDataSize(sizeDelta int64) {
	d.dataSize += sizeDelta
}

func (d Data) Size() int64 {
	return d.dataSize
}
