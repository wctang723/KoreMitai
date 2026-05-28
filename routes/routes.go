package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wctang723/KoreMitai/middleware"
)

func SetTimeoutRoutes(router *gin.Engine) {
	router.Use(middleware.TimeoutMiddleware(10 * time.Second))
}
