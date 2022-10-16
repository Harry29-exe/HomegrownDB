package math2

import "HomegrownDB/common/datastructs"

// Power returns a to power of b
func Power[T datastructs.Number](a T, b int) (result T) {
	result = a
	for i := 0; i < b-1; i++ {
		result *= a
	}

	return
}
