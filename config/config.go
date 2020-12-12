package config

import (
	"os"
	"strconv"
)

// Conf - Struct for all configurations
type Conf struct {
	*Gin
	GinServer      GinServer
	PostgresConfig PostgresConfig
}

// PostgresConfig - Initializes DB server details
type PostgresConfig struct {
	DBInstance      string
	DBName          string
	DBServer        string
	DBPort          int
	DBUser          string
	DBPwd           string
	DBRetryAttempts uint8
}

// GinServer - Initializes server details
type GinServer struct {
	bindIPAddress string
	bindPort      int
}

// Init - Initialize application configuration properties
func Init() *Conf {
	var conf Conf

	conf.GinServer.bindIPAddress = os.Getenv("GO_GIN_IP_ADDR")
	conf.GinServer.bindPort, _ = strconv.Atoi(os.Getenv("GO_GIN_IP_PORT"))

	conf.PostgresConfig.DBInstance = os.Getenv("GO_DB_INSTANCE")
	conf.PostgresConfig.DBName = os.Getenv("GO_DB_NAME")
	conf.PostgresConfig.DBServer = os.Getenv("GO_DB_IP_ADDR")
	conf.PostgresConfig.DBPort, _ = strconv.Atoi(os.Getenv("GO_DB_IP_PORT"))
	conf.PostgresConfig.DBUser = os.Getenv("GO_DB_USER")
	conf.PostgresConfig.DBPwd = os.Getenv("GO_DB_PWD")
	conf.PostgresConfig.DBRetryAttempts = 2

	return &conf
}
