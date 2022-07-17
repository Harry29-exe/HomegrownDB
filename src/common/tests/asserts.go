package tests

import "testing"

func AssertEq[T comparable](v1, v2 T, t *testing.T) {
	if v1 != v2 {
		t.Errorf("Value %+v and %+v are not equal", v1, v2)
	}
}

func AssertEqArray[T comparable](v1, v2 []T, t *testing.T) {
	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			t.Errorf("Values at intex %d: %+v and %+v are not equal", i, v1, v2)
		}
	}
}
