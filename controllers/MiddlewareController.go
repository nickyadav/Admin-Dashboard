package controllers

import (
	"strings"

	"github.com/astaxie/beego/context"
)

var FilterUser = func(c *context.Context) {
	//Setting important Header
	c.Output.Header("Cache-Control", "no-cache,no-store,max-age=0,must-revalidate")
	c.Output.Header("Pragma", "no-cache")
	c.Output.Header("Expires", "-1")
	c.Output.Header("X-XSS-Protection", "1;mode=block")
	c.Output.Header("X-Frame-Options", "SAMEORIGIN")
	c.Output.Header("X-Content-Type-Options", "nosniff")
	if strings.HasPrefix(c.Input.URL(), "/account/login") {
		return
	}
	_, ok := c.Input.Session("GlobalSession").(string)
	if !ok {
		if strings.HasPrefix(c.Input.URL(), "/account/logout") {
			c.Redirect(302, "/account/login")
		} else {
			c.Redirect(302, "/account/login?r="+c.Input.URL())
		}
	}
}
