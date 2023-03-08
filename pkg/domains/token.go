package domain

type IToken interface {
	GenerateJWT(email, username string) (string, error)
	ValidateToken(signedToken string) error
}
