package random

import (
	"math/rand"
	"strings"
)

type Random struct {
	*rand.Rand
}

func (r Random) StringASCII(length uint) string {
	builder := strings.Builder{}
	for i := uint(0); i < length; i++ {
		builder.WriteRune(r.CharASCII())
	}

	return builder.String()
}

func (r Random) StringASCIIRandLen(minLen, maxLen uint) string {
	length := r.Int64(int64(minLen), int64(maxLen))
	return r.StringASCII(uint(length))
}

func (r Random) CharASCII() rune {
	return rune(r.Int64(32, 126))
}

// Int64 returns value in range <min, max>
func (r Random) Int64(min, max int64) int64 {
	return (r.Int63() + min) % max
}
