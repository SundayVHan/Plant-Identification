package main

import (
	"github.com/gin-gonic/gin"
	"plant_identification/config"
	"plant_identification/internal/article"
	"plant_identification/internal/database"
	"plant_identification/internal/identification"
	"plant_identification/internal/router"
	"plant_identification/internal/thumb"
	"plant_identification/internal/user"
)

func main() {
	// 初始化 config
	config.LoadConfig()

	database.Init()
	user.Init()
	identification.Init()
	thumb.Init()
	article.Init()

	r := gin.Default()
	router.RegisterRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
