package router

import (
	"github.com/gin-gonic/gin"
	"plant_identification/internal/article"
	"plant_identification/internal/identification"
	"plant_identification/internal/thumb"
	"plant_identification/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	userApi := api.Group("/user")
	userApi.GET("/create", user.Register)
	userApi.GET("/login", user.Login)

	identApi := api.Group("/identification")
	identApi.Use(AuthMiddleware())
	identApi.POST("/qa", identification.QA)
	identApi.GET("/history", identification.GetHistory)
	identApi.GET("/star", identification.SetStar)

	thumbApi := api.Group("/thumb")
	thumbApi.Use(AuthMiddleware())
	thumbApi.GET("/add", thumb.AddThumb)
	thumbApi.GET("/sum", thumb.CountThumbs)

	articleApi := api.Group("/article")
	articleApi.Use(AuthMiddleware())
	articleApi.GET("/publish", article.PublishArticle)
	articleApi.GET("/fetch", article.FetchArticle)
	articleApi.POST("/generation", article.GenerateArticle)
}
