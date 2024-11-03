package router

import (
	"github.com/gin-gonic/gin"
	"plant_identification/internal/identification"
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
}
