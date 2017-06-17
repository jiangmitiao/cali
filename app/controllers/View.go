package controllers

import "github.com/revel/revel"

type View struct {
	*revel.Controller
}

func (c View) Index() revel.Result {
	return c.RenderTemplate("View/cali.html")
}

func (c View) Public() revel.Result {
	return c.RenderTemplate("View/public.html")
}
func (c View) Book() revel.Result {
	return c.RenderTemplate("View/book.html")
}
func (c View) Login() revel.Result {
	return c.RenderTemplate("View/login.html")
}
func (c View) SignUp() revel.Result {
	return c.RenderTemplate("View/signup.html")
}
func (c View) Person() revel.Result {
	return c.RenderTemplate("View/person.html")
}

func (c View) Read() revel.Result {
	return c.RenderTemplate("View/read.html")
}
