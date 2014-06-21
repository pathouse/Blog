package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var h = &helpers.Helper{}
	return c.Render(h)
}
