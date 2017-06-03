package services

import (
	"github.com/jiangmitiao/cali/app/models"
)

//all tags count
func QueryTagsCount() int64 {
	count, _ := engine.Count(models.Tag{})
	return count
}

//tags info
func QueryTags(limit, start int) []models.Tag {
	tags := make([]models.Tag, 0)
	engine.Limit(limit, start).Find(&tags)
	return tags
}
