package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wctang723/KoreMitai/config"
)

func GetAnimes(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()

		animesID := ctx.Param("animesID")
		animesUUID, err := uuid.Parse(animesID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if animesID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No anime id present!"})
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
