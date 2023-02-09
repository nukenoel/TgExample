package main

import (
	"TelegramBot/config"
	"TelegramBot/internal/bot"
	"TelegramBot/pkg/postgresql"
	"context"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.GetConfig()
	postgresqlClient, err := postgresql.NewClient(context.TODO(), cfg)
	if err != nil {
		log.Fatalf("unable to create postgresql client, descripiton:%v", err)
	}
	app, err := bot.NewApp(cfg, postgresqlClient)
	if err != nil {
		log.Error(err)
		return
	}
	app.Run()
	log.Info("running telegram bot application")
}
