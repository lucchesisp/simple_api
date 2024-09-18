package services

import (
	"github.com/lucchesisp/simple_api/src/interfaces"
)

type UserService struct {
	repository interfaces.UserRepository
}

func NewUserService(repository interfaces.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) GetUserById(id string) string {
	// Valida se o usuario eh valido
	// manda log pra serviço de auditoria
	// manda email pro usuario

	user := u.repository.FindById(id)

	if user == "" {
		return "User not found"
	}

	return user
}
