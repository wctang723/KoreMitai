package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wctang723/KoreMitai/config"
)

func GetAnime(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		animesID := ctx.Param("animesid")

		animesUUID, err := uuid.Parse(animesID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Parsing went wrong", "error_msg": err.Error()})
			return
		}

		animesInfo, err := cfg.Myqu.GetAnimes(c, animesUUID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No anime found!"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"animes_id": animesInfo.AnimesID, "animes_title": animesInfo.Title})
	}
}

func GetAnimes(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
