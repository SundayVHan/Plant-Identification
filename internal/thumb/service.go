package thumb

import "plant_identification/internal/database"

func Init() {
	err := database.DB.AutoMigrate(&Thumb{})
	if err != nil {
		panic(err)
	}
}
