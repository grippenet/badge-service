package main

import(
	"os"
	"fmt"
	"log/slog"
	"github.com/grippenet/badge-service/pkg/config"
	"github.com/grippenet/badge-service/pkg/db"
	"github.com/grippenet/badge-service/pkg/db/memory"
	"github.com/grippenet/badge-service/pkg/services"
	"github.com/grippenet/badge-service/pkg/server"
	"github.com/grippenet/badge-service/pkg/types"
)

func main() {

	appConfig, err := config.LoadConfig()
	if(err != nil) {
		slog.Error(fmt.Sprintf("Error loading config : %s\n", err))
		os.Exit(1)
	}

	var dbService types.DBService

	if(appConfig.DBConfig.URI == config.MemoryDbURI) {
		dbService = memory.NewMemoryDBService()
	} else {
		dbService, err = db.NewBadgeDBService(appConfig.DBConfig)
		if(err != nil) {
			slog.Error(fmt.Sprintf("Unable to connect to db : %s\n", err))
			os.Exit(1)
		}
	}

	svc := services.InitServices(dbService)

	httpServer := server.NewHttpServer(appConfig.Http, svc)

	err = httpServer.Start()

	if(err != nil) {
		slog.Error(fmt.Sprintf("Unable to launch server : %s\n", err))
		os.Exit(1)
	}

}