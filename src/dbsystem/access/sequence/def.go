package sequence

import "HomegrownDB/lib/datastructs"

type Sequence[T datastructs.Number] interface {
	Next() T
}
