package app

import (
	"github.com/gin-gonic/gin"
	"plant_identification/config"
	"plant_identification/internal/database"
	"plant_identification/internal/router"
)

func main() {
	// 初始化 config
	config.LoadConfig()

	database.Init()

	r := gin.Default()
	router.RegisterRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
