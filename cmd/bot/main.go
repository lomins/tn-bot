package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/tn-bot/internal/config"
	"github.com/tn-bot/internal/storage"
	"github.com/tn-bot/internal/telegram"
)

func main() {
	token, ok := os.LookupEnv("TELEGRAM_APITOKEN")
	if !ok {
		log.Fatalf("Can't find .env or token")
	}

	spreadsheetIdAdmin, ok := os.LookupEnv("SPREADSHEET_ID_ADMIN")
	if !ok {
		log.Fatalf("Can't find .env or spreadsheetId")
	}

	spreadsheetIdUsers, ok := os.LookupEnv("SPREADSHEET_ID_USERS")
	if !ok {
		log.Fatalf("Can't find .env or spreadsheetId")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	cfg := config.New()

	pg := storage.NewPostgres(cfg)
	defer pg.Close()

	sheetsAdmin := storage.NewSheets(spreadsheetIdAdmin, "Admin")

	sheetsUsers := storage.NewSheets(spreadsheetIdUsers, "User")

	telegramBot := telegram.NewBot(bot, pg, sheetsAdmin, sheetsUsers)
	telegramBot.Start()
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
