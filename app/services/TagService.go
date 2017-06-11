package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"strconv"
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

//query more tags by bookid
func QueryBookTags(bookid int) []models.Tag {
	tags := make([]models.Tag, 0)
	engine.SQL("select tags.* from tags,books_tags_link where tags.id=books_tags_link.tag and book=" + strconv.Itoa(bookid)).Find(&tags)
	return tags
}
