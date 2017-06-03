package controllers

import "github.com/revel/revel"

type Install struct {
	*revel.Controller
}

func (c Install) Index() revel.Result {
	return c.RenderTemplate("admin/install/index.html")
}
