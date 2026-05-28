package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	// "github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wctang723/KoreMitai/api"
	"github.com/wctang723/KoreMitai/config"
	"github.com/wctang723/KoreMitai/database"
	"github.com/wctang723/KoreMitai/routes"
)

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	defer db.Close()

	var apiCfg config.ApiConfig
	apiCfg.Platform = os.Getenv("PLATFORM")
	apiCfg.Myqu = dbQueries
	apiCfg.Tokensecretkey = os.Getenv("JWTTOKENSECRET")

	myrouter := gin.Default()
	myrouter.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	myrouter.GET("/animes", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	myrouter.GET("/films", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	myrouter.POST("/register/create", api.UserRegister(&apiCfg))
	myrouter.POST("/login", api.UserLogin(&apiCfg))
	myrouter.GET("/reviews/:reviewsid", api.GetReviews(&apiCfg))
	myrouter.GET("/animes/:animesid", api.GetAnimes(&apiCfg))

	// TODO: routes package not implemented yet
	routes.SetTimeoutRoutes(myrouter)

	myrouter.Run(":8080")
	// log.Fatal(autotls.Run(myrouter, "localhost:8080"))
}
