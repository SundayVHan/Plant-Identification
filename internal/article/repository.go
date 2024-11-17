package article

import "plant_identification/internal/database"

func addArticle(text string, imgBase64 string, title string) error {
	article := Article{
		Text:      text,
		ImgBase64: imgBase64,
		Title:     title,
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func getArticles() ([]Article, error) {
	var articles []Article
	if err := database.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
