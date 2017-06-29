package services

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/google/uuid"
	"time"
)

type CaliCategoryService struct {
	
}

func (service CaliCategoryService) QueryCount()int  {
	count,_:=engine.Count(models.CaliCategory{})
	return int(count)
}

func (service CaliCategoryService) Query()[]models.CaliCategory  {
	categories := make([]models.CaliCategory,0)
	engine.Find(&categories)
	return categories
}

func (service CaliCategoryService) GetOrInsertCategoryByName(categoryName string)( category models.CaliCategory)  {
	if ok, _ := engine.Where("category = ?", categoryName).Get(&category); ok {
		return
	} else {
		category.Id = uuid.New().String()
		category.Category = categoryName
		category.CreatedAt = time.Now().Unix()
		category.UpdatedAt = time.Now().Unix()
		engine.InsertOne(category)
		return
	}
}