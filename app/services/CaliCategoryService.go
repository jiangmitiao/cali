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
	if categoryId == models.DefaultCaliCategory.Id {
		return
	}
	engine.Where("id = ?", categoryId).Delete(models.CaliCategory{})
	engine.Where("cali_category = ?", categoryId).Delete(models.CaliBookCategory{})
}

func (service CaliCategoryService) DeleteBookCategoryByBookId(bookId string) {
	engine.Where("cali_book = ?", bookId).Delete(models.CaliBookCategory{})
}

func (service CaliCategoryService) QueryByBookIdWithOutDefault(bookid string) (categories []models.CaliCategory) {
	engine.Where("id in  (select cali_category from cali_book_category where cali_book = ?)", bookid).And("id != ?", "default").Find(&categories)
	if len(categories) == 0 {
		engine.Where("id = ?", "default").Find(&categories)
	}
	return
}
