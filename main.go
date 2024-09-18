package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/simple_api/src/controller"
	"github.com/lucchesisp/simple_api/src/repository"
	"github.com/lucchesisp/simple_api/src/services"
)

func main() {
	engine := gin.Default()

	databaseConnection := "jdba=localhost user=postgres dbname=postgres sslmode=disable"

	userRepository := repository.NewUserRepositoryLocal(databaseConnection)

	userService := services.NewUserService(&userRepository)
	userController := controller.NewUserController(userService)

	engine.GET("/user/:id", userController.GetUserById)

	engine.Run(":8080")
}

// REQUEST -> CONTROLLER -> SERVICE (regra de negocio) -> REPOSITORY(conversa com o banco)
