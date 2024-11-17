package article

import "plant_identification/internal/database"

func Init() {
	err := database.DB.AutoMigrate(&Article{})
	if err != nil {
		panic(err)
	}
}
