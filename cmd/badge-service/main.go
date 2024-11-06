package main

import(
	"os"
	"fmt"
	"log/slog"
	"github.com/grippenet/badge-service/pkg/config"
	"github.com/grippenet/badge-service/pkg/db"
	"github.com/grippenet/badge-service/pkg/services"
	"github.com/grippenet/badge-service/pkg/server"
)

func main() {

	appConfig, err := config.LoadConfig()
	if(err != nil) {
		slog.Error(fmt.Sprintf("Error loading config : %s\n", err))
		os.Exit(1)
	}

	dbService, err := db.NewBadgeDBService(appConfig.DBConfig)
	if(err != nil) {
		slog.Error(fmt.Sprintf("Unable to connect to db : %s\n", err))
		os.Exit(1)
	}

	svc := services.InitServices(dbService)

	httpServer := server.NewHttpServer(appConfig.Http, svc)

	err = httpServer.Start()

	if(err != nil) {
		slog.Error(fmt.Sprintf("Unable to launch server : %s\n", err))
		os.Exit(1)
	}

}