package main

import (
	"database/sql"
	"flag"
	"os"

	"github.com/VoC925/go-testify/internal"
	"github.com/VoC925/go-testify/internal/api"
	"github.com/VoC925/go-testify/internal/api/config"
	"github.com/VoC925/go-testify/internal/user"
	"github.com/VoC925/go-testify/internal/user/db"
	"github.com/VoC925/go-testify/pkg/logging"
)

var (
	// путь до файла конфига
	configPathFl = flag.String("pathCfg", "config.yml", "path to config file")
	dbPathFl     = flag.String("pathDB", "user.db", "path to DB file")
)

func main() {
	flag.Parse()

	// логгер
	logger := logging.New()

	// config
	cfg := config.New(*configPathFl)
	logger.Trace("config registered")

	// клиент для ДБ
	dbUser, err := sql.Open(cfg.Storage.Name, *dbPathFl)
	if err != nil {
		logger.Fatal(internal.ErrPingConn)
		return
	}
	if err := dbUser.Ping(); err != nil {
		logger.Fatal(internal.ErrPingConn)
		return
	}
	defer dbUser.Close()

	logger.Trace("client registered")

	// хранилище
	store := db.New(dbUser)
	logger.Trace("store registered")
	// сервис
	service := user.NewService(store, logger)
	logger.Trace("service registered")

	// хендлеры
	handlerService := user.NewHandlerUser(service, logger)

	// Создание экземпляра структура Server
	s := api.New(cfg, handlerService)
	// Запуск сервера
	if err := s.Start(); err != nil {
		logger.Errorf("server: %s", err)
		os.Exit(1)
	}
}
