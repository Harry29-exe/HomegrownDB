package auth

type Ctx interface {
}

type Manager interface {
	Authenticate(auth Authentication) (User, error)
}
