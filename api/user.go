package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wctang723/KoreMitai/auth"
	"github.com/wctang723/KoreMitai/config"
	"github.com/wctang723/KoreMitai/database"
	"github.com/wctang723/KoreMitai/model"
)

// WARNING: Don't return the real err message to the client!(right now just use for conviencies) Remember to change it later!

func UserRegister(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.CreateUserForm
		c := ctx.Request.Context()

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// TODO: Might need to add some email validation
		hashedpasswd, err := auth.HashPassword(user.Passwd)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		user_params := database.CreateUserParams{
			UserID:         user.UserID,
			Email:          user.Email,
			HashedPassword: hashedpasswd,
		}

		if _, err := cfg.Myqu.CreateUser(c, user_params); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user_id": user.UserID})
	}
}

func UserLogin(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.LoginUserForm
		c := ctx.Request.Context()

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var userinfo database.User

		if user.Email == "" || user.Passwd == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing email or password!"})
			return
		}

		userinfo, err := cfg.Myqu.GetUserByEmail(c, user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		passwdcheck, err := auth.CheckPasswordHash(user.Passwd, userinfo.HashedPassword)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if passwdcheck != true {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email or password"})
			return
		}

		userJWT, err := auth.MakeJWT(userinfo.ID, cfg.Tokensecretkey, time.Hour)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userRefreshToken := auth.MakeRefreshToken()
		createRefreshTokenParams := database.CreateRefreshTokenParams{
			Token:  userRefreshToken,
			UserID: uuid.NullUUID{UUID: userinfo.ID},
		}

		rt, err := cfg.Myqu.CreateRefreshToken(c, createRefreshTokenParams)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"refresh_token": rt.Token, "jwt": userJWT})
	}
}
