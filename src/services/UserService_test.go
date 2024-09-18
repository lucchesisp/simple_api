package services

import (
	"github.com/go-playground/assert/v2"
	"github.com/lucchesisp/simple_api/src/repository"
	"testing"
)

func TestUserService(t *testing.T) {

	userRepository := repository.NewUserRepositoryLocal("")

	t.Run("GetUserByIdGabriel", func(t *testing.T) {
		userService := NewUserService(&userRepository)
		result := userService.GetUserById("14444")

		assert.Equal(t, result, "Gabriel")
	})

	t.Run("GetUserByIdNotFound", func(t *testing.T) {
		userService := NewUserService(&userRepository)
		result := userService.GetUserById("123")

		assert.Equal(t, result, "User not found")
	})
}
