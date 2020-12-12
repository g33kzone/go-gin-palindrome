package route

import (
	"github.com/g33kzone/go-gin-palindrome/db"
	"github.com/gin-gonic/gin"
)

// GinRoutes - Initialize Routes
type GinRoutes struct {
	GinEngine *gin.Engine
}

// InitializeRoutes - Initialize application route
func (routes GinRoutes) InitializeRoutes(db *db.Db) {
	routerGroup := routes.GinEngine.Group("g33kzone/v1")
	initRoute(routerGroup, db)
}
