package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)

		log.Println(example)
	})

	router := gin.Default()
	router.Use(ErrorHandler())

	router.GET("/ok", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	router.GET("/error", func(ctx *gin.Context) {
		ctx.Error(errors.New("something went wrong"))
	})

	router2 := gin.Default()
	authorized := router2.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":     "bar",
		"austion": "1234",
		"lena":    "hello2",
		"manu":    "4321",
	}))

	authorized.GET("/secrets", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	router2.Run(":8080")

	// router.Run(":8088")

	// r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Set("example", "12345")

		ctx.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := ctx.Writer.Status()
		log.Println(status)
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
