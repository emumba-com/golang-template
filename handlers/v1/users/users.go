package users

import (
	"github.com/gofrs/uuid"
	appLogger "gogintemplate/logger"
	"gogintemplate/models/v1"
	"gogintemplate/services/db"
	"gogintemplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store       db.Store
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

// CreateUser creates a new user
// @Summary Create User
// @Description Create a new user
// @Tags users
// @Produce  json
// @Success 200 {object} models.Users
// @Failure 400	{object} models.Response
// @Failure 500	{object} models.Response
// @Router /users/ [post].
func (server *Server) CreateUser(ctx *gin.Context) {
	logger := appLogger.GetLogger()

	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		logger.Error(err.Error())
		utils.BuildResponse(ctx, http.StatusBadRequest, utils.ERROR, err.Error(), nil)

		return
	}

	newUser, err := server.Store.CreateUser(user)
	if err != nil {
		logger.Error(err.Error())
		statusCode, errMsg := utils.ParseDBError(err, "Users")
		utils.BuildResponse(ctx, statusCode, utils.ERROR, errMsg, nil)

		return
	}

	utils.BuildResponse(ctx, http.StatusOK, utils.SUCCESS, "", newUser)
}

// GetUser
// @Summary Get user by id
// @Description Get user by user id
// @Tags users
// @Produce  json
// @Success 200 {object} models.Users
// @Failure 400	{object} models.Response
// @Failure 500	{object} models.Response
// @Router /users/{id}/ [get].
func (server *Server) GetUser(ctx *gin.Context) {
	logger := appLogger.GetLogger()

	userID, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		logger.Error(err.Error())
		utils.BuildResponse(ctx, http.StatusBadRequest, utils.ERROR, err.Error(), nil)

		return
	}

	user, err := server.Store.GetUser(userID)
	if err != nil {
		logger.Error(err.Error())
		statusCode, errMsg := utils.ParseDBError(err, "Users")
		utils.BuildResponse(ctx, statusCode, utils.ERROR, errMsg, nil)

		return
	}

	utils.BuildResponse(ctx, http.StatusOK, utils.SUCCESS, "", user)
}
