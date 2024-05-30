package ports

type ForAuthorizingUser interface {
	Authorize(username string) (bool, error)
}
