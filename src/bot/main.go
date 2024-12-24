package main

import (
	"database/sql"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		log.Fatalf("TELEGRAM_API_KEY not set in .env file")
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		// Copy empty.db to data.db if there's an error opening data.db
		input, err := os.ReadFile("./db/empty.db")
		if err != nil {
			log.Fatalf("Error reading empty.db file: %v", err)
		}

		err = os.WriteFile("./db/data.db", input, 0644)
		if err != nil {
			log.Fatalf("Error writing data.db file: %v", err)
		}

		// Try opening data.db again
		db, err = sql.Open("sqlite3", "./db/data.db")
		if err != nil {
			log.Fatalf("Error opening data.db file: %v", err)
		}
	}
	defer db.Close()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
