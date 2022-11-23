package tokenizer

import (
	"sort"
)

type charCollection interface {
	In(char rune) bool
}

func newCharCollection(slice []rune) charCollection {
	if len(slice) > 30 {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		return &sortedCharCollection{slice, len(slice)}
	} else {
		return smCharCollection{slice}
	}
}

type sortedCharCollection struct {
	sortedSlice []rune
	len         int
}

func (c *sortedCharCollection) In(char rune) bool {
	return sort.Search(c.len,
		func(i int) bool {
			return c.sortedSlice[i] >= char
		}) < c.len
}

type smCharCollection struct {
	slice []rune
}

func (c smCharCollection) In(char rune) bool {
	for _, ch := range c.slice {
		if ch == char {
			return true
		}
	}
	return false
}
