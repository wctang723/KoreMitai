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

	// TODO: test all the routes handler work as expected!
	myrouter.POST("/register", api.UserRegister(&apiCfg))
	myrouter.POST("/login", api.UserLogin(&apiCfg))

	myrouter.GET("/animes", api.GetAnimes(&apiCfg))
	myrouter.GET("/reviews", api.GetReviews(&apiCfg))

	myrouter.GET("/animes/:animesid", api.GetAnime(&apiCfg))
	myrouter.GET("/reviews/:reviewsid", api.GetReview(&apiCfg))

	// NOTE: routes package not implemented yet
	routes.SetTimeoutRoutes(myrouter)

	myrouter.Run(":8080")
	// log.Fatal(autotls.Run(myrouter, "localhost:8080"))
}
