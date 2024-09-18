package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/simple_api/src/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{userService}
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	user := u.userService.GetUserById(ctx.GetString("id"))
	
	ctx.JSON(200, gin.H{"user found:": user})
}
