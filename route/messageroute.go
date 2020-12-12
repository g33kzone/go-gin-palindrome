package route

import (
	"github.com/g33kzone/go-message-palindrome/db"
	"github.com/g33kzone/go-message-palindrome/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func initRoute(msgRoutes *gin.RouterGroup, db *db.Db) {
	msgRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	msgRoutes.GET("/notes/:id", func(c *gin.Context) {
		handler.FetchMessageHandler(c, db)
	})
	msgRoutes.GET("/notes", func(c *gin.Context) {
		handler.FetchMessageListHandler(c, db)
	})
	msgRoutes.POST("/notes", func(c *gin.Context) {
		handler.SaveMessageHandler(c, db)
	})
	msgRoutes.DELETE("/notes/:id", func(c *gin.Context) {
		handler.DeleteMessageHandler(c, db)
	})
}
