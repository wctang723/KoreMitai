package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wctang723/KoreMitai/config"
)

func GetReview(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		reviewID := ctx.Param("reviewsid")

		reviewUUID, err := uuid.Parse(reviewID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reviewInfo, err := cfg.Myqu.GetReviews(c, reviewUUID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"review_id": reviewInfo.AnimesID, "review_body": reviewInfo.Body.String, "review_star": reviewInfo.Star})
	}
}

func GetReviews(cfg *config.ApiConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
