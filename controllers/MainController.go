package controllers

type MainController struct {
	baseRouter
}

// @router / [get]
func (c *MainController) Index() {
	c.TplName = "dashboard.html"
}
