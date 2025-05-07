package app

import (
	api "github.com/balticoldschool/user-service/.generated"
	"github.com/balticoldschool/user-service/src/domain/user"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApp() {
	api.RegisterHandlers(router, user.NewUserHandler())
	log.Fatal(router.Run(":8080"))
}
