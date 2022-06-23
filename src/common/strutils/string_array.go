package strutils

import "fmt"

type StrArray struct {
	Array []string
}

func (sa *StrArray) FormatAndAdd(format string, args ...any) {
	sa.Array = append(sa.Array, fmt.Sprintf(format, args...))
}

func (sa *StrArray) Add(str string) {
	sa.Array = append(sa.Array, str)
}
