package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/wctang723/KoreMitai/internal/database"
)

type apiConfig struct {
	myqu           *database.Queries
	platform       string
	tokensecretkey string
}

type User struct {
	ID        uuid.UUID `json:"id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Passwd    string    `json:"password" binding:"required"`
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	defer db.Close()

	var apiCfg apiConfig
	apiCfg.platform = os.Getenv("PLATFORM")
	apiCfg.myqu = dbQueries
	apiCfg.tokensecretkey = os.Getenv("JWTTOKENSECRET")

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	router.GET("/animes", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	router.GET("/films", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	router.POST("/users/create", apiCfg.UserRegister())

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	s.ListenAndServe()
}

func (cfg *apiConfig) UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User
		c := ctx.Request.Context()

		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user_params := database.CreateUserParams{
			UserID: user.UserID,
			Email:  user.Email,
		}

		if _, err := cfg.myqu.CreateUser(c, user_params); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"user_id": user.UserID})
	}
}
