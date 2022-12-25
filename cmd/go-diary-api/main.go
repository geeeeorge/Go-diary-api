package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/geeeeorge/Go-diary-api/config"
	"github.com/geeeeorge/Go-diary-api/infra/database"
	"github.com/geeeeorge/Go-diary-api/infra/server"
)

func main() {
	conf := config.Load()
	db, err := database.NewPostgresDB(conf.RDBConfig.ConnectionString())
	if err != nil {
		log.Fatalf(err.Error())
	}
	s := server.NewServer(conf.AppPort(), conf.AppHost(), db, nil)
	s.Start()
}
