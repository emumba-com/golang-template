package main

import (
	"errors"
	"gogintemplate/docs"
	"gogintemplate/env"
	"gogintemplate/logger"
	"gogintemplate/middlewares"
	"gogintemplate/services/db"
	"net/http"

	"time"

	"gogintemplate/controllers/v1/health"
	"gogintemplate/controllers/v1/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("Starting Gin Service")
	setupSwaggerDocumentation()

	database := db.GetConnection()
	dbStore := db.NewStore(database)

	router := gin.New()

	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middlewares.LoggerMiddleware())

	ginServiceGrp := router.Group("gin-service/api/v1")

	health.CreateNewServer(dbStore, router, ginServiceGrp)
	users.CreateNewServer(dbStore, router, ginServiceGrp)

	// register swagger documentation endpoint
	ginServiceGrp.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	httpServer := &http.Server{
		Addr:    ":" + env.Env.ServerPort,
		Handler: router,
	}

	startHttpServer(httpServer)
}

func setupSwaggerDocumentation() {
	docs.SwaggerInfo.Title = "Emumba - Gin Framework API Documentation"
	docs.SwaggerInfo.Description = "This swagger documentation contains API documentation for Gin Framework Boilerplate" +
		"REST endpoints."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/gin-service/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func startHttpServer(server *http.Server) {
	logger := logger.GetLogger()
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error())
	}
}
