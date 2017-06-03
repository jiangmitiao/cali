package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
)

type Language struct {
	*revel.Controller
}

func (c Language) Index(callback string) revel.Result {
	return c.RenderJSONP(callback, models.NewOKApi())
}

//all languages count
func (c Language) LanguagesCount(callback string) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryLanguagesCount()))
}

//all languages info
func (c Language) Languages(callback string, limit, start int) revel.Result {
	return c.RenderJSONP(
		callback,
		models.NewOKApiWithInfo(services.QueryLanguages(limit, start)),
	)
}
