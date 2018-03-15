package handler

import "github.com/gin-gonic/gin"

type pingHandler struct {
}

func (pingHandler) Ping(c *gin.Context) {
	c.HTML(200, "ping.tmpl", gin.H{
		"msg": "con cac",
	})
}
