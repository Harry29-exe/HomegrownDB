package inputtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	"fmt"
)

func ConvStr(val string) ([]byte, error) {
	l := len(val) + 4
	if l > hgtype.MaxVarLenSize {
		return nil, fmt.Errorf("string values can not be longer that %d", hgtype.MaxVarLenSize) //todo better err
	}
	serializedVal := make([]byte, l)
	bparse.Serialize.PutUInt4(uint32(l), serializedVal)
	copy(serializedVal[4:], val)
	return serializedVal, nil
}
