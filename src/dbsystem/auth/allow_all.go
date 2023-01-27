package auth

func NewAllowAllManager() Manager {
	return allowAllManager{}
}

var _ Manager = allowAllManager{}

type allowAllManager struct{}

func (a allowAllManager) Authenticate(auth Authentication) (User, error) {
	return NewAnonymous(), nil
}
