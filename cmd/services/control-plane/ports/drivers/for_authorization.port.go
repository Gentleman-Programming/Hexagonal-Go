package ports

type ForManagingAuthorization interface {
	Authorize(username string) (bool, error)
}
