package handler

import (
	"net/http"

	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/params/request"
	"example.com/gin_realworld/params/response"
	"example.com/gin_realworld/security"
	"example.com/gin_realworld/utils"
	"github.com/gin-gonic/gin"
)

func AddUserHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/users")
	usersGroup.POST("", userRegistration)
	usersGroup.POST("/login", userLogin)
}

func userRegistration(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserRegistrationRequest
	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		log.WithError(err).Errorf("user registration bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithField("user", utils.JsonMarshal(body)).Infof("user registration called")

	//TODO: insert data to db
	token, err := security.GenerateJWT(body.User.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{User: response.UserAuthenticationBody{
		Email:    body.User.Email,
		Token:    token,
		Username: body.User.Username,
		Bio:      "",
		Image:    "https://api.realworld.io/images/smiley-curus.jpg",
	}})
	return
}

func userLogin(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserLoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		log.WithError(err).Errorf("user login bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.WithField("user", utils.JsonMarshal(body)).Infof("user login called")
	//TODO: get username from DB
	userName := "xxx123456"
	token, err := security.GenerateJWT(userName, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{User: response.UserAuthenticationBody{
		Email:    body.User.Email,
		Token:    token,
		Username: userName,
		Bio:      "",
		Image:    "https://api.realworld.io/images/smiley-curus.jpg",
	}})
	return
}
