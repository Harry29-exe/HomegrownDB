package random

import (
	"math/rand"
	"strings"
	"time"
)

type Random struct {
	*rand.Rand
}

func NewRandomByTime() Random {
	return Random{
		rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func NewRandom(seed int64) Random {
	return Random{rand.New(rand.NewSource(seed))}
}

func (r Random) StringASCII(length uint) string {
	builder := strings.Builder{}
	for i := uint(0); i < length; i++ {
		builder.WriteRune(r.CharASCII())
	}

	return builder.String()
}

func (r Random) StringASCIIRandLen(minLen, maxLen uint) string {
	length := r.Int64mm(int64(minLen), int64(maxLen))
	return r.StringASCII(uint(length))
}

func (r Random) CharASCII() rune {
	return rune(r.Int64mm(32, 126))
}

// Int64mm returns value in range <min, max>
func (r Random) Int64mm(min, max int64) int64 {
	return (r.Int63() + min) % max
}

func (r Random) Int16() int16 {
	return int16(r.Rand.Int31())
}

func (r Random) Int16mm(min, max int16) int16 {
	return (int16(r.Int31()) - min) % max
}
