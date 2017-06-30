package api

import (
	"github.com/revel/revel"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"strings"
)

type Category struct {
	*revel.Controller
}


func (c Category) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c Category)All()revel.Result  {
	return c.RenderJSONP(c.Request.FormValue("callback"),models.NewOKApiWithInfo(categoryService.Query()))
}

func (c Category) Add()revel.Result  {
	categoryName :=rcali.ValueOrDefault(strings.Trim(c.Request.FormValue("categoryName")," "),models.DefaultCaliCategory.Category)
	categoryService.GetOrInsertCategoryByName(categoryName)
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c Category) Update()revel.Result  {
	categoryName :=strings.Trim(c.Request.FormValue("categoryName")," ")
	categoryId := strings.Trim(c.Request.FormValue("categoryId")," ")
	categoryService.UpdateCategoryName(categoryId,categoryName)
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

func (c Category) Delete()revel.Result  {
	categoryId := strings.Trim(c.Request.FormValue("categoryId")," ")
	categoryService.DeleteById(categoryId)
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}
