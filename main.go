package main

import (
	"flag"
	"os"

	"github.com/g33kzone/go-message-palindrome/config"
	"github.com/g33kzone/go-message-palindrome/db"
	"github.com/g33kzone/go-message-palindrome/docs"
	"github.com/g33kzone/go-message-palindrome/route"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	profile := flag.String("profile", "dev", "Development profile")
	if *profile == "dev" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.TraceLevel)
		log.SetOutput(os.Stdout)
		log.SetReportCaller(false)
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
		log.SetLevel(log.TraceLevel)
		log.SetOutput(os.Stdout)
		log.SetReportCaller(false)
	}
}

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Palindrome Validator - Notes API"
	docs.SwaggerInfo.Description = "Validator Service - by g33kzone."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "pragmatic-video-app.herokuapp.com"
	docs.SwaggerInfo.BasePath = "/g33kzone/v1"
	// docs.SwaggerInfo.Schemes = []string{"https"}

	conf := config.Init()
	ginServer := config.Gin{GinEngine: gin.Default(), ServerConfig: conf.GinServer, PostgresConfig: conf.PostgresConfig}

	db := &db.Db{}
	db.InitDBConnection(conf.PostgresConfig, 1)

	routes := route.GinRoutes{GinEngine: ginServer.GinEngine}
	routes.InitializeRoutes(db)

	ginServer.StartServer()
}
