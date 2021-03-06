package controller

import (
	"echoException/excel-faker/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Favicon(c *gin.Context) {
	path := "resource/image/favicon.ico"
	c.File(path)
}

func HTMLhandler(c *gin.Context) {
	page := c.Param("page")
	if page == "" {
		page = "index"
	}
	page = page + ".html"
	c.HTML(http.StatusOK, page, gin.H{
		"APP_URL": helper.GetAppUrl(c),
	})
}
