package services

import "github.com/jiangmitiao/cali/app/models"

//all languages count
func QueryLanguagesCount() int64 {
	count, _ := engine.Count(models.Language{})
	return count
}

//languages info
func QueryLanguages(limit, start int) []models.Language {
	languages := make([]models.Language, 0)
	engine.Limit(limit, start).Find(&languages)
	return languages
}
