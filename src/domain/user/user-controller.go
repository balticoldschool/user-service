package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserHandler interface {
	CreateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context, userId uuid.UUID)
}

type controller struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return controller{
		service,
	}
}

func (c controller) CreateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "hallo welt")
}

func (c controller) GetUserById(ctx *gin.Context, userId uuid.UUID) {
	user, err := c.service.FindUserById(userId)
	if err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
