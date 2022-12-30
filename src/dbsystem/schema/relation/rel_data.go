package relation

import "HomegrownDB/common/bparse"

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

func (d Data) serialize() []byte {
	serializer := bparse.NewSerializer()
	serializer.Int64(d.dataSize)
	serializer.MdString(d.dirFilepath)

	return serializer.GetBytes()
}

func deserializeData(deserializer *bparse.Deserializer) data {
	return data{
		dataSize:    deserializer.Int64(),
		dirFilepath: deserializer.MdString(),
	}
}
