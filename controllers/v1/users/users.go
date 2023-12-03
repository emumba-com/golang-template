package users

import (
	"github.com/gin-gonic/gin"
	"gogintemplate/handlers/v1/users"
	"gogintemplate/services/db"
)

func registerUserRoutes(server *users.Server) {
	usersRoutes := server.RouterGroup.Group("users")
	{
		usersRoutes.POST("/", server.CreateUser)
	}

	usersRoutes = server.RouterGroup.Group("users/:id")
	{
		usersRoutes.GET("/", server.GetUser)
	}
}

func CreateNewServer(dbStore db.Store, router *gin.Engine, rg *gin.RouterGroup) {
	server := &users.Server{
		Store:       dbStore,
		Router:      router,
		RouterGroup: rg,
	}
	registerUserRoutes(server)
}
