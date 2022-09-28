package column

import "HomegrownDB/dbsystem/ctype"

type Args interface {
	Type() (ctype.Type, error)
	Name() (string, error)
	Nullable() (bool, error)
	Length() (uint32, error)
}
