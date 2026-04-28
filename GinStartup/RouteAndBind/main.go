package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// type Person struct {
// 	Name     string    `form:"name"`
// 	Address  string    `form:"address"`
// 	Birthday time.Time `form:"birthday" time_format:"DateOnly" time_utc:"1"`
// }

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(200, message)
	})

	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")
		ctx.String(200, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "anonymous")

		ctx.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/post", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page := ctx.DefaultQuery("page", "0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		ctx.String(200, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// router.POST("/post", func(ctx *gin.Context) {
	// 	ids := ctx.QueryMap("ids")
	// 	names := ctx.PostFormMap("names")
	//
	// 	fmt.Printf("ids: %v; names: %v\n", ids, names)
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"ids":   ids,
	// 		"names": names,
	// 	})
	// })

	// external redirect
	router.GET("/old", func(ctx *gin.Context) {
		ctx.Redirect(301, "https://www.google.com/")
	})
	router.HEAD("/old", func(ctx *gin.Context) {
		ctx.Redirect(301, "https://www.google.com/")
	})

	// redirect from POST
	router.POST("/submit", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/result")
	})

	// internal router redirect
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/final"
		router.HandleContext(ctx)
	})

	router.GET("/final", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})

	router.GET("result", func(ctx *gin.Context) {
		ctx.String(200, "redirect here!")
	})

	router.GET("/api/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		// simulate a lookup
		if id == "0" {
			Fail(ctx, http.StatusNotFound, "USER_NOT_FOUND", "no user with that ID")
			return
		}
		OK(ctx, gin.H{"id": id, "name": "Alice"})
	})

	// router.Use(VersionMiddleware())
	// router.GET("/api/users", func(ctx *gin.Context) {
	// 	version := ctx.GetString("api_version")
	//
	// 	switch version {
	// 	case "v2":
	// 		ctx.JSON(200, gin.H{"version": "v2", "data": []gin.H{}})
	// 	default:
	// 		ctx.JSON(200, gin.H{"version": "v1", "users": []string{}})
	// 	}
	// })

	router.POST("loginJSON", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "you are loggin in "})
	})

	// router.GET("/testing", startPage)
	// router.POST("/testing", startPage)
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		h := testHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"reat": h.Rate, "Domain": h.Domain})
	})

	r.Run(":8086")

	route.Run(":8088")

	router.Run(":8080")
}

func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(ctx *gin.Context, status int, code, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}

// VersionMiddleware reads the API version from the Accept-Version header. This keeps the URLs clean but the client to set the custom header
func VersionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		version := ctx.GetHeader("Accept-Version")
		if version == "" {
			version = "v1"
		}
		ctx.Set("api_version", version)
		ctx.Next()
	}
}

// func startPage(c *gin.Context) {
// 	var person Person
//
// 	// NOTE: There are some issue with ShouldBind method when it comes to the JSON data.
// 	// The time parsing would not work as expected as set with time_format.
// 	// The issue detail: https://grok.com/share/c2hhcmQtMg_ad94e8ed-1e9f-4f53-b954-b82b5415ce03 & Gin issues #1193
// 	if err := c.ShouldBind(&person); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

//
// 	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person.Name, person.Address, person.Birthday)
// 	c.JSON(http.StatusOK, gin.H{
// 		"name":     person.Name,
// 		"address":  person.Address,
// 		"birthday": person.Birthday,
// 	})
// }
