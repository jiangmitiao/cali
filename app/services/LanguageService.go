package services

import "github.com/jiangmitiao/cali/app/models"

type LanguageService struct {
}

//all languages count
func (service LanguageService) QueryLanguagesCount() int64 {
	count, _ := engine.Count(models.Language{})
	return count
}

//languages info
func (service LanguageService) QueryLanguages(limit, start int) []models.Language {
	languages := make([]models.Language, 0)
	engine.Limit(limit, start).Find(&languages)
	return languages
}
