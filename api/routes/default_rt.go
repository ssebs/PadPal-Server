package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/util"
)

// RootHandler renders the REST-API.md file as HTML for /
func RootHandler(c *gin.Context) {
	// get contents of ./REST-API.md and return
	md, err := os.ReadFile("./REST-API.md")
	if err != nil {
		c.Error(err)
		c.JSON(500, err)
		return
	}

	// convert to html
	html := util.ParseMDToHTML(md)
	c.Header("content-type", "text/html")
	c.String(200, string(html))
}

// ErrorHandler will log the error and return JSON with the message
func ErrorHandler(code int, err error, c *gin.Context) {
	c.Error(err)
	c.JSON(code, err.Error())
}
