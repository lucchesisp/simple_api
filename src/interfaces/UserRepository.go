package interfaces

type UserRepository interface {
	FindById(id string) string
}
