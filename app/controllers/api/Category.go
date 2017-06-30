package api

import (
	"github.com/revel/revel"
	"github.com/jiangmitiao/cali/app/models"
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
	return nil
}
