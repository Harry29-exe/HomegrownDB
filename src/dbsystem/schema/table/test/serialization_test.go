package table_test

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"reflect"
	"testing"
)

func TestSerializeSimpleTable(t *testing.T) {
	//given
	def := table.NewDefinition("users")

	err := addColumns(def,
		column.NewDefinition("id", false, hgtype.NewInt8(hgtype.Args{})),
		column.NewDefinition("name", false, hgtype.NewStr(hgtype.Args{Length: 255})),
	)
	if err != nil {
		return
	}

	//when
	serializer := bparse.NewSerializer()
	def.Serialize(serializer)
	deserializedDef := table.Deserialize(serializer.GetBytes())
	//then
	assert.True(reflect.DeepEqual(def, deserializedDef), t)
}

func addColumns(definition table.Definition, columns ...column.WDef) error {
	for _, colDef := range columns {
		err := definition.AddColumn(colDef)
		if err != nil {
			return err
		}
	}
	return nil
}
