package domain

type AuthSvc interface {
	GenerateToken(id string, isAdmin bool) (string, error)
}
