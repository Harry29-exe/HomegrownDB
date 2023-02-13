package bparse

type Serializable interface {
	Serialize(s *Serializer)
	Deserialize(d *Deserializer)
}
