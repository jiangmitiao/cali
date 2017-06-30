package services

import (
	"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"time"
)

type CaliCategoryService struct {
}

func (service CaliCategoryService) QueryCount() int {
	count, _ := engine.Count(models.CaliCategory{})
	return int(count)
}

func (service CaliCategoryService) Query() []models.CaliCategory {
	categories := make([]models.CaliCategory, 0)
	engine.Find(&categories)
	return categories
}

func (service CaliCategoryService) GetOrInsertCategoryByName(categoryName string) (category models.CaliCategory) {
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

func (service CaliCategoryService) UpdateCategoryName(categoryId, categoryName string) {
	engine.Where("id = ?", categoryId).Cols("category", "updated").Update(models.CaliCategory{Id: categoryId, Category: categoryName, UpdatedAt: time.Now().Unix()})
}

func (service CaliCategoryService) DeleteById(categoryId string) {
	engine.Where("id = ?", categoryId).Delete(models.CaliCategory{})
}
