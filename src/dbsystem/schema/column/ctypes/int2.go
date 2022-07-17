package ctypes

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/column"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
)

const Int2 column.Type = "Int2"

func NewInt2Column(args column.Args) *Int2Column {
	name, err := args.Name()
	if err != nil {
		panic("No column name")
	}
	nullable, _ := args.Nullable()

	return &Int2Column{
		name:       name,
		nullable:   nullable,
		parser:     &int2Parser{nullable},
		serializer: &int2Serializer{nullable},
	}
}

type Int2Column struct {
	prototype
	name     string
	nullable bool

	parser     *int2Parser
	serializer *int2Serializer
}

func (c *Int2Column) Name() string {
	return c.name
}

func (c *Int2Column) Nullable() bool {
	return c.nullable
}

func (c *Int2Column) Type() column.Type {
	return Int2
}

func (c *Int2Column) DataParser() column.DataParser {
	return c.parser
}

func (c *Int2Column) DataSerializer() column.DataSerializer {
	return c.serializer
}

func (c *Int2Column) Serialize() []byte {
	serializer := bparse.NewSerializer()
	serializer.MdString(Int2)
	serializer.MdString(c.name)
	serializer.Bool(c.nullable)

	return serializer.GetBytes()
}

func (c *Int2Column) Deserialize(data []byte) []byte {
	deserializer := bparse.NewDeserializer(data)
	//skip Column Type
	_ = deserializer.MdString()
	c.name = deserializer.MdString()
	c.nullable = deserializer.Bool()
	c.serializer = &int2Serializer{columnIsNullable: c.nullable}
	c.parser = &int2Parser{columnIsNullable: c.nullable}

	return deserializer.RemainedData()
}

type int2Parser struct {
	columnIsNullable bool
}

func (i *int2Parser) Skip(data []byte) []byte {
	return data[:2]
}

func (i *int2Parser) Parse(data []byte) (column.Value, []byte) {
	v, next := bparse.Deserialize.Int2(data)
	value := NewInt2Value(&v)

	return value, next
}

type int2Serializer struct {
	columnIsNullable bool
}

func (s *int2Serializer) SerializeValue(value any) (column.DataToSave, error) {
	switch data := value.(type) {
	case int:
		if data < math.MaxInt16 {
			return column.NewDataToSaveInTuple(bparse.Serialize.Int2(int16(data))), nil
		} else {
			return nil, errors.New(fmt.Sprintf(
				"can save integer %d bigger that %d as int2", data, math.MaxInt16))
		}
	case int16:
		return column.NewDataToSaveInTuple(bparse.Serialize.Int2(data)), nil
	}

	return nil, errors.New(fmt.Sprintf(
		"value: %+v, is nor int nor int16 type", value))
}

func (s *int2Serializer) SerializeInput(value string) (column.DataToSave, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return column.NewDataToSave(bparse.Serialize.Uint2(uint16(v)), column.StoreInTuple), nil
}

type int2Value struct {
	value        *int16
	valueAsBytes []byte
}

func NewInt2Value(value *int16) *int2Value {
	if value == nil {
		return &int2Value{
			value:        nil,
			valueAsBytes: nil,
		}
	}

	asBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(asBytes, uint16(*value))

	return &int2Value{
		value:        value,
		valueAsBytes: asBytes,
	}
}

func (v *int2Value) AsBytes() []byte {
	asBytes := make([]byte, 0, 2)
	binary.LittleEndian.PutUint16(asBytes, uint16(*v.value))

	return asBytes
}

func (v *int2Value) Value() any {
	return v.value
}

func (v *int2Value) IsNull() bool {
	return v.value == nil
}

func (v *int2Value) EqualsBytes(value []byte) bool {
	if v.value == nil {
		if value == nil {
			return true
		}
		return false
	} else if value == nil {
		return false
	}

	if len(value) != 2 {
		panic("Value should be exactly 2 bytes long")
	}
	return bytes.Compare(v.valueAsBytes, value) == 0
}

func (v *int2Value) Equals(value *any) bool {
	if v.value == nil {
		if value == nil {
			return true
		}
		return false
	} else if value == nil {
		return false
	}

	switch data := (*value).(type) {
	case int16:
		return *v.value == data
	default:
		panic("TYPE of argument and column value are different")
	}
}

func (v *int2Value) CompareBytes(value []byte) int {
	if v.value == nil {
		if value == nil {
			return 0
		}
		return -1
	} else if value == nil {
		return 1
	}

	if len(value) != 2 {
		panic("Value should be exactly 2 bytes long")
	}
	return bytes.Compare(v.valueAsBytes, value)
}

func (v *int2Value) Compare(value *any) int {
	if v.value == nil {
		if value == nil {
			return 0
		}
		return -1
	} else if value == nil {
		return 1
	}

	switch data := (*value).(type) {
	case int16:
		if *v.value < data {
			return -1
		} else if *v.value > data {
			return 1
		} else {
			return 0
		}
	default:
		panic("TYPE of argument value and column value are different")
	}
}

func (v *int2Value) ToString() string {
	return strconv.Itoa(int(*v.value))
}
