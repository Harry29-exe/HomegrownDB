package assert

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func Eq[T comparable](v1, v2 T, t *testing.T) {
	if v1 != v2 {
		t.Errorf("Value %+v and %+v are not equal", v1, v2)
		debug.PrintStack()
	}
}

func NotEq[T comparable](v1, v2 T, t *testing.T) {
	if v1 == v2 {
		t.Errorf("Value %+v and %+v are equal", v1, v2)
		debug.PrintStack()
	}
}

func EqDeep[T any](v1, v2 T, t *testing.T) {
	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("Value %+v and %+v are not equal", v1, v2)
		debug.PrintStack()
	}
}

func True(value bool, t *testing.T) {
	if !value {
		t.Errorf("Given value was expected to be true")
		debug.PrintStack()
	}
}

func NotNil(val interface{}, t *testing.T) {
	if val == nil {
		t.Errorf("Value is nil")
		debug.PrintStack()
	}
}

func IsNil(val interface{}, t *testing.T) {
	if val != nil {
		t.Errorf("Value: %#v is not nil", val)
		debug.PrintStack()
	}
}

func ErrNotNil(val error, t *testing.T) {
	if val == nil {
		t.Errorf("Assert failed, expected err got nil")
		debug.PrintStack()
	}
}

func ErrIsNil(val error, t *testing.T) {
	if val != nil {
		t.Errorf("Assert failed, expected nil got err: %s", val.Error())
		debug.PrintStack()
	}
}

func EqArray[T comparable](v1, v2 []T, t *testing.T) {
	if len(v1) != len(v2) {
		t.Errorf("Arrays have different lenghts")
		debug.PrintStack()
		return
	}
	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			t.Errorf("Values at intex %d: %+v and %+v are not equal", i, v1[i], v2[i])
			debug.PrintStack()
			return
		}
	}
}
