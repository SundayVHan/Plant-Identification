package router

import (
	"github.com/gin-gonic/gin"
	"plant_identification/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	userApi := api.Group("/user")
	userApi.POST("/create", user.Register)
	userApi.GET("/login", user.Login)

	identApi := api.Group("/identification")
	identApi.Use(AuthMiddleware())
	identApi.GET("/qa")
	identApi.GET("history")
}
