package buffer_test

import (
	"HomegrownDB/dbsystem/io/buffer"
	"HomegrownDB/tests/testutils"
	"testing"
)

var testTables = testutils.TestTables

func TestSharedBuffer_Overflow(t *testing.T) {
	testBuffer := buffer.NewSharedBuffer(10_000)
}
