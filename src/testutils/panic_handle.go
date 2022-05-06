package testutils

import "testing"

func ShouldPanic(fun func(), noPanicMsg string, t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error(noPanicMsg)
		}
	}()

	fun()
}
