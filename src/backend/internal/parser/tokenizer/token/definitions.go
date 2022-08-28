package token

import (
	"strconv"
)

type Token interface {
	Code() Code
	Value() string
}

func NewBasicToken(code Code, value string) Token {
	return &BasicToken{
		code:  code,
		value: value,
	}
}

type BasicToken struct {
	code  Code
	value string
}

func (b *BasicToken) Code() Code {
	return b.code
}

func (b *BasicToken) Value() string {
	return b.value
}

func NilToken() Token {
	return &BasicToken{
		code:  Nil,
		value: "",
	}
}

func NewIntegerToken(value string) (*IntegerToken, error) {
	integer, err := strconv.Atoi(value)
	if err != nil {
		return nil, err
	}

	return &IntegerToken{
		Token: NewBasicToken(Integer, value),
		Int:   integer,
	}, nil
}

type IntegerToken struct {
	Token
	Int int
}

func NewFloatToken(value string) (*FloatToken, error) {
	float, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	return &FloatToken{
		Token: NewBasicToken(Integer, value),
		Float: float,
	}, nil
}

type FloatToken struct {
	Token
	Float float64
}
