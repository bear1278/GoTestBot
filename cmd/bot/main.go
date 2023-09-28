package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/bear1278/GoTestBot/pkg/config"
	Mysql "github.com/bear1278/GoTestBot/pkg/db/MySQL"
	telegram "github.com/bear1278/GoTestBot/pkg/telegram"
)

func main() {

	cfg, err := config.Init()

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	db, err := sql.Open("mysql", cfg.DbPtah)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	DbRepository := Mysql.NewDbRepository(db)

	telegramBot := telegram.NewBot(bot, DbRepository, cfg.Messages)

	if err = telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
