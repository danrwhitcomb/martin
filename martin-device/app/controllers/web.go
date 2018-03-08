package controllers

import (
	"github.com/revel/revel"
)

type WebController struct {
	*revel.Controller
}

func (c WebController) Index() revel.Result {
	return c.Render()
}
