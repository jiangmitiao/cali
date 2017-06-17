package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"strconv"
)

type Language struct {
	*revel.Controller
}

func (c Language) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

//all languages count
func (c Language) LanguagesCount() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(languageService.QueryLanguagesCount()))
}

//all languages info
func (c Language) Languages() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(languageService.QueryLanguages(limit, start)),
	)
}
