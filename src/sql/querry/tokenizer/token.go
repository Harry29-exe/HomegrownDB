package tokenizer

type Token interface {
	Code() TokenCode
	Value() string
}

func CreateToken(code TokenCode, value string) Token {
	return &basicToken{
		code:  code,
		value: value,
	}
}

type TokenCode = uint16

const (
	Select TokenCode = iota
	From
	Where
	Join
	Update
	Delete
	CreateTable
	DropTable
)

type basicToken struct {
	code  TokenCode
	value string
}

func (b *basicToken) Code() TokenCode {
	return b.code
}

func (b *basicToken) Value() string {
	return b.value
}
