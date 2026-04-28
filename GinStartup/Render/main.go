package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	router := gin.Default()

	// router.LoadHTMLGlob("templates/*")
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/posts/index", func(ctx *gin.Context) {
		ctx.HTML(200, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/users/index", func(ctx *gin.Context) {
		ctx.HTML(200, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.GET("/user", func(ctx *gin.Context) {
		user := gin.H{"name": "Lena", "role": "admin"}

		switch ctx.Query("format") {
		case "xml":
			ctx.XML(200, user)
		case "yaml":
			ctx.YAML(200, user)
		default:
			ctx.JSON(200, user)
		}
	})

	router.GET("/someJSON", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(ctx *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		ctx.JSON(200, msg)
	})

	router.GET("/someXML", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/someProtoBuf", func(ctx *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		ctx.ProtoBuf(200, data)
	})

	// Standard JSON
	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Pure JSON
	router.GET("/purejson", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.Run(":8080")
}
