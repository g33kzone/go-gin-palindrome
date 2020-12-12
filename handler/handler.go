package handler

import (
	"github.com/gin-gonic/gin"
)

//setCors handles cross origin requests, api requests via a web browser to an api on a different domain
func setCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Authorization, Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, If-Modified-Since, X-File-Name, Cache-Control")
	}
}

//globalRecover returns json if there is an unhandled exception
// func globalRecover(c *gin.Context) {
// 	defer func(c *gin.Context) {
// 		if rec := recover(); rec != nil {
// 			c.JSON(http.StatusInternalServerError, model.APIMessage{Code: 500, Text: fmt.Sprintf("Internal Server Error, %v", rec)})
// 		}
// 	}(c)
// 	c.Next()
// }
