package services

import "github.com/jiangmitiao/cali/app/models"

//all authors count
func QueryAuthorsCount() int64 {
	count, _ := engine.Count(models.Author{})
	return count
}

//authors info
func QueryAuthors(limit, start int) []models.Author {
	authors := make([]models.Author, 0)
	engine.Limit(limit, start).Find(&authors)
	return authors
}
