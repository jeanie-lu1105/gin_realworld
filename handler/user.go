package handler

import (
	"net/http"
	"time"

	"example.com/gin_realworld/cache"
	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/middlewares"
	"example.com/gin_realworld/models"
	"example.com/gin_realworld/params/request"
	"example.com/gin_realworld/params/response"
	"example.com/gin_realworld/security"
	"example.com/gin_realworld/storage"
	"example.com/gin_realworld/utils"
	"github.com/bsm/redislock"
	"github.com/gin-gonic/gin"
)

func AddUserHandler(r *gin.Engine) {
	usersGroup := r.Group("/api/users")
	usersGroup.POST("", userRegistration)
	usersGroup.POST("/login", userLogin)
	usersGroup.POST("/user-image", uploadUserImage)
	r.GET("/api/profiles/:username", userProfile)
	r.Use(middlewares.AuthMiddlewareCookie).PUT("/api/user", editUser)
}

func userRegistration(ctx *gin.Context) {
	log := logger.New(ctx)
	var body request.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.WithError(err).Errorln("bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.WithField("user", utils.JsonMarshal(body)).Infof("user registration called")

	defaultUserImage := "https://api.realworld.io/images/smiley-cyrus.jpeg"

	hashPassword, err := security.HashPassword(body.User.Password)
	if err != nil {
		log.WithError(err).Errorf("hash password failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := storage.CreateUser(ctx, &models.User{
		Username: body.User.Username,
		Password: hashPassword,
		Email:    body.User.Email,
		Image:    defaultUserImage,
		Bio:      "",
	}); err != nil {
		log.WithError(err).Errorf("create user failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := security.GenerateJWT(body.User.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.SetCookie("token", token, 24*3600, "/", "localhost:8080", false, false)

	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{
		User: response.UserAuthenticationBody{
			Email:    body.User.Email,
			Token:    token,
			Username: body.User.Username,
			Bio:      "",
			Image:    defaultUserImage,
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

	dbUser, err := storage.GetUserByEmail(ctx, body.User.Email)
	log.Infof("db user info %s\n", utils.JsonMarshal(dbUser))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !security.ValidatePassword(body.User.Password, dbUser.Password) {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := security.GenerateJWT(dbUser.Username, body.User.Email)
	if err != nil {
		log.WithError(err).Errorln("generate jwt failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{User: response.UserAuthenticationBody{
		Email:    body.User.Email,
		Token:    token,
		Username: dbUser.Username,
		Bio:      dbUser.Bio,
		Image:    dbUser.Image,
	}})
}

func userProfile(ctx *gin.Context) {
	log := logger.New(ctx)
	userName := ctx.Param("username")
	log = log.WithField("username", userName)
	log.Infof("user profile called, userName: %v", userName)
	user, err := storage.GetUserByUsername(ctx, userName)
	if err != nil {
		log.WithError(err).Errorln("get user by username failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := cache.SetUserProfile(ctx, security.GetCurrentUsername(ctx), user, 300); err != nil {
		log.WithError(err).Errorln("set user profile failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, response.UserProfileResponse{
		UserProfile: response.UserProfile{Username: user.Username,
			Bio: user.Bio, Image: user.Image, Following: false},
	})
}

func editUser(ctx *gin.Context) {
	log := logger.New(ctx)
	log.Infof("edit user: %v", security.GetCurrentUsername(ctx))

	lock, err := cache.Locker.Obtain(ctx, cache.UserEditProfileLockKey(security.
		GetCurrentUsername(ctx)), 30*time.Second, &redislock.Options{
		RetryStrategy: redislock.ExponentialBackoff(100*time.Millisecond, 2*time.Second),
	})

	if err == redislock.ErrNotObtained {
		log.WithError(err).Errorln("user is editing" + security.GetCurrentUsername(ctx))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if err != nil {
		log.WithError(err).Errorln("obtain lock failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer lock.Release(ctx)

	time.Sleep(5 * time.Second)

	var body request.EditUserRequest
	if err := ctx.BindJSON(&body); err != nil {
		log.WithError(err).Errorf("edit user bind json failed")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if body.EditUserBody.Username == "" || body.EditUserBody.Email == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if body.EditUserBody.Password != "" {
		var err error
		body.EditUserBody.Password, err = security.HashPassword(body.EditUserBody.Password)
		if err != nil {
			log.WithError(err).Errorln("hash password failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		dbUser := &models.User{
			Username: body.EditUserBody.Username,
			Password: body.EditUserBody.Password,
			Email:    body.EditUserBody.Email,
			Image:    body.EditUserBody.Image,
			Bio:      body.EditUserBody.Bio,
		}

		if err := storage.UpdateUserByUsername(ctx, security.GetCurrentUsername(ctx), dbUser); err != nil {
			log.WithError(err).Errorln("update user failed")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := cache.DeleteUserProfile(ctx, security.GetCurrentUsername(ctx)); err != nil {
			log.WithError(err).Errorln("delete user profile failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		token, err := security.GenerateJWT(dbUser.Username, body.EditUserBody.Email)
		if err != nil {
			log.WithError(err).Errorln("generate jwt failed")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, response.UserAuthenticationResponse{User: response.UserAuthenticationBody{
			Email:    dbUser.Email,
			Token:    token,
			Username: dbUser.Username,
			Bio:      dbUser.Bio,
			Image:    dbUser.Image,
		}})
	}
}

func uploadUserImage(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	err := ctx.SaveUploadedFile(file, "./"+file.Filename)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}
