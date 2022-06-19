package ctypes_test

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	factory2 "HomegrownDB/dbsystem/schema/column/factory"
	"reflect"
	"testing"
)

func TestInt2Column_Builder(t *testing.T) {
	nullable := false
	name := "small_int"

	args := column.ArgsBuilder().
		Type(ctypes.Int2).Nullable(nullable).
		Name(name).
		Build()

	col := factory2.CreateDefinition(args)
	if col.Nullable() != nullable {
		t.Errorf("nullable param. expected: %t, actual: %t",
			nullable, col.Nullable())
	}
	if col.Name() != name {
		t.Errorf("name param. expected: %s, actual: %s",
			name, col.Name())
	}
}

func TestInt2Column_Serializable(t *testing.T) {
	name := "my_int_with2_bytes"
	columnDef := createInt2(name, false)

	serialized := columnDef.Serialize()
	deserialized, _ := factory2.DeserializeColumnDefinition(serialized)
	recreatedDef, ok := deserialized.(*ctypes.Int2Column)
	if !ok {
		t.Errorf("Can not cast result")
	}
	equals := reflect.DeepEqual(*columnDef, *recreatedDef)
	if !equals {
		t.Errorf("Column definitions does not match. "+
			"Expected: %#v, Actual: %#v", columnDef, recreatedDef)
	}
}

func createInt2(name string, nullable bool) *ctypes.Int2Column {
	args := column.ArgsBuilder().
		Type(ctypes.Int2).Nullable(nullable).
		Name(name).
		Build()

	return ctypes.NewInt2Column(args)
}
