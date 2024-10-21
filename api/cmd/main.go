package main

import (
	"context"
	"fmt"
	"my_wallet/api/server"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	dir, err := os.Getwd()
	if err != nil {
		logrus.Panic("Error finding the work directory:", err)
	}
	fmt.Println("Work directory actually:", dir)
	entries, err := os.ReadDir("./")
	if err != nil {
		logrus.Fatal(err)
	}

	for _, e := range entries {
		logrus.Info(e.Name())
	}

	envPath := "C:/Users/miguel.gn/Documents/Practica/go/wallet/My-Wallet/.env"
	//envPath := filepath.Join(dir, ".env")  //For container
	logrus.Infof("find file.env in: %s", envPath)

	viper.SetConfigFile(envPath)

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("Error al leer el archivo de configuración:", err)
	}
	httpAddr := viper.GetString("SERVER_PORT_HTTP")
	dburl := viper.GetString("DB_URL")
	srv, err := server.New(logger, httpAddr, dburl, ctx)
	if err != nil {
		logger.Panic("Failed to create server:", err)
	}

	srv.Start()

}
