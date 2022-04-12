package table

var Deserializer = tableDeserializer{}

type tableDeserializer struct{}

func (d tableDeserializer) Deserialize(data []byte) Definition {
	table := &table{}
	table.Deserialize(data)

	return table
}
