package main

import (
	"context"
	"my_wallet/api/server"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title My Wallet API
// @version 1.0
// @description This is a sample server for a wallet API.
// @host localhost:8081
// @BasePath /
func main() {
	ctx := context.Background()
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	dir, err := os.Getwd()
	if err != nil {
		logrus.Panic("Layer: main ", "Error: finding the work directory:", err)
	}
	logger.Infoln("Layer: main ", "Message: Work directory actually:", dir)
	entries, err := os.ReadDir("./")
	if err != nil {
		logrus.Fatal(err)
	}

	for _, e := range entries {
		logrus.Info(e.Name())
	}
	//rootDir := filepath.Join(dir, "../..")
	envPath := filepath.Join(dir, ".env") //For container replace rootDir for dir and for local use rootDir
	logrus.Infoln("Layer: main ", "Message: find file.env in: %s", envPath)
	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("Layer: main ", "Error al leer el archivo de configuración:", err)
	}
	httpAddr := viper.GetString("SERVER_PORT_HTTP")
	dburl := viper.GetString("DB_URL")
	srv, err := server.New(logger, httpAddr, dburl, ctx)
	if err != nil {
		logger.Panic("Layer: main ", "Failed to create server:", err)
	}

	srv.Start()

}
