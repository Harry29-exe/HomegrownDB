package auth

type Authentication interface {
	Type() MethodType
}

type MethodType int8

const (
	UsernameAndPassw MethodType = iota
	Token
	None
)
