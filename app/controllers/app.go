package controllers

import (
	"Blog/app/helpers"
	"Blog/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var plink, ulink helpers.Linker
	h := &helpers.Helper{}
	plink = &models.Post{}
	ulink = &models.User{}
	return c.Render(h, plink, ulink)
}
