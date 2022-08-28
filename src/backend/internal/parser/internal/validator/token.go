package validator

import (
	"HomegrownDB/backend/internal/parser/internal/source"
	"HomegrownDB/backend/internal/parser/sqlerr"
	"HomegrownDB/backend/internal/parser/tokenizer/token"
)

type tokenValidator struct {
	source source.TokenSource
	token  token.Token
	err    error
}

func Next(source source.TokenSource) *tokenValidator {
	source.Checkpoint()
	return &tokenValidator{
		source: source,
		token:  source.Next(),
		err:    nil,
	}
}

func NextSequence(source source.TokenSource, codes ...token.Code) error {
	source.Checkpoint()
	for _, code := range codes {
		next := source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), source)
			source.Rollback()
			return err
		}
	}

	source.Commit()
	return nil
}

func Current(source source.TokenSource) *tokenValidator {
	source.Checkpoint()
	return &tokenValidator{
		source: source,
		token:  source.Current(),
		err:    nil,
	}
}

func CurrentSequence(source source.TokenSource, codes ...token.Code) error {
	currentToken := source.Current()
	if currentToken.Code() != codes[0] {
		return sqlerr.NewTokenSyntaxError(codes[0], currentToken.Code(), source)
	}

	source.Checkpoint()
	for _, code := range codes[1:] {
		next := source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), source)
			source.Rollback()
			return err
		}
	}

	source.Commit()
	return nil
}

func (tc *tokenValidator) Check() (token.Token, error) {
	if tc.err != nil {
		tc.source.Rollback()
		return nil, tc.err
	}
	tc.source.Commit()
	return tc.token, nil
}

func (tc *tokenValidator) Has(code token.Code) *tokenValidator {
	switch {
	case tc.err != nil:
		break

	case tc.token == nil:
		tc.err = sqlerr.NewSyntaxError(token.ToString(code), "nil", tc.source)

	case tc.token.Code() != code:
		tc.err = sqlerr.NewSyntaxError(token.ToString(code), tc.token.Value(), tc.source)
	}

	return tc
}

func (tc *tokenValidator) IsTextToken() *textTokenChecker {
	tc.Has(token.Identifier)
	if tc.err != nil {
		return nilTextTokenChecker(tc)
	}

	switch textToken := tc.token.(type) {
	case *token.TextToken:
		return &textTokenChecker{
			tokenValidator: tc,
			textToken:      textToken,
		}
	default:
		tc.err = sqlerr.NewSyntaxError("Identifier", tc.token.Value(), tc.source)
		return nilTextTokenChecker(tc)
	}
}
