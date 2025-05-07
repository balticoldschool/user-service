package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler interface {
	CreateUser(ctx *gin.Context)
}

type controller struct{}

func NewUserHandler() UserHandler {
	return controller{}
}

func (c controller) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "hallo welt")
}
