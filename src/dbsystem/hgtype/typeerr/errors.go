package typeerr

type ToLongErr struct{}

func (t ToLongErr) Error() string {
	return "data to long"
}

type UTF8NotAllowed struct{}

func (u UTF8NotAllowed) Error() string {
	return "utf8 not allowed"
}

type NullNotAllowed struct{}

func (n NullNotAllowed) Error() string {
	return "null not allowed"
}

type TypesNotConvertable struct{}

func (t TypesNotConvertable) Error() string {
	return "type is not assignable"
}
