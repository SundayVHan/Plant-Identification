package thumb

import "plant_identification/internal/database"

func addThumb(userId int64) error {
	thumb := Thumb{
		UserId: userId,
	}

	if err := database.DB.Create(&thumb).Error; err != nil {
		return err
	}

	return nil
}

func countThumbs() (int64, error) {
	var count int64
	if err := database.DB.Model(&Thumb{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
