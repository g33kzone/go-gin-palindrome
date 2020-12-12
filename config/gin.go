package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Gin - Gin configurations Struct
type Gin struct {
	GinEngine      *gin.Engine
	ServerConfig   GinServer
	PostgresConfig PostgresConfig
}

// StartServer - Initialize Gin server struct
func (gin Gin) StartServer() {
	log.Infof("Attempting to start Gin server - %s:%d",
		gin.ServerConfig.bindIPAddress, gin.ServerConfig.bindPort)

	// db := &db.Db{}
	// db = &Db{}
	// db.InitDBConnection(gin.PostgresConfig, 1)

	// routes := route.GinRoutes{GinEngine: gin.GinEngine}
	// routes.InitializeRoutes()

	err := gin.GinEngine.Run(fmt.Sprintf("%s:%d",
		gin.ServerConfig.bindIPAddress, gin.ServerConfig.bindPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start Gin server %s", err.Error()))
	}
}
