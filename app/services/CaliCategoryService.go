package services

import "github.com/jiangmitiao/cali/app/models"

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