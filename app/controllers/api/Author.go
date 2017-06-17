package api

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/revel/revel"
	"strconv"
)

type Author struct {
	*revel.Controller
}

func (c Author) Index() revel.Result {
	return c.RenderJSONP(c.Request.FormValue("callback"), models.NewOKApi())
}

//all tags count
func (c Author) AuthorsCount() revel.Result {
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(authorService.QueryAuthorsCount()))
}

//all tags info
func (c Author) Authors() revel.Result {
	limit, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("limit"), rcali.ClassNumsStr))
	start, _ := strconv.Atoi(rcali.ValueOrDefault(c.Request.FormValue("start"), "0"))
	return c.RenderJSONP(
		c.Request.FormValue("callback"),
		models.NewOKApiWithInfo(authorService.QueryAuthors(limit, start)),
	)
}
