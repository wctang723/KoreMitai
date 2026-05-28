package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wctang723/KoreMitai/config"
)

func GetReviews(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		reviewID := ctx.Param("reviewsID")
		reviewUUID, err := uuid.Parse(reviewID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if reviewID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No review id present!"})
			return
		}

		reviewInfo, err := cfg.Myqu.GetReviews(c, reviewUUID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"review_id": reviewInfo.AnimesID, "review_body": reviewInfo.Body, "review_star": reviewInfo.Star})
	}
}
