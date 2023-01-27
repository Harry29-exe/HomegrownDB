package auth

type User interface {
	Name() string
	Anonymous() bool
}

func NewUser(name string) User {
	return &StdUser{
		Username:    name,
		IsAnonymous: false,
	}
}

func NewAnonymous() User {
	return &StdUser{
		Username:    "",
		IsAnonymous: true,
	}
}

var _ User = &StdUser{}

type StdUser struct {
	Username string

	IsAnonymous bool
}

func (s *StdUser) Name() string {
	return s.Username
}

func (s *StdUser) Anonymous() bool {
	return s.IsAnonymous
}


