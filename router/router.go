package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wctang723/KoreMitai/auth"
	"github.com/wctang723/KoreMitai/config"
	"github.com/wctang723/KoreMitai/database"
	"github.com/wctang723/KoreMitai/model"
)

func UserRegister(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		c := ctx.Request.Context()

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

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
