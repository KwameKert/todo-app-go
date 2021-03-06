package app

import (
	//"encoding/json"

	// "strconv"
	"gorm.io/gorm"

	"todo/app/models"
	"todo/app/repository"
	"todo/app/routes"
	"todo/app/services"
	"todo/core"
	"todo/core/database"

	log "github.com/sirupsen/logrus"
)

type App struct{}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func (app *App) Start(conf *core.Config) {
	log.Info("Starting Todolist API server")
	pg := setupDatabase(conf)
	repo := repository.NewRepository(pg)
	services := services.NewService(repo, conf)

	server := core.NewHTTPServer(conf)
	router := routes.NewRouter(server.Engine, conf, services)

	router.RegisterRoutes()
	server.Start()

}

func setupDatabase(conf *core.Config) *gorm.DB {
	pg, err := database.NewPostgres(conf)
	if err != nil {
		log.Fatal("failed to initialize postgres database. err:", err)
		panic(err.Error())
	}
	err = database.RunMigrations(pg, &models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("failed to run migrations. err:", err)
	}

	if conf.RUN_SEEDS {
		models.RunSeeds(pg)
	}
	return pg
}
