package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"leakGuard/parse"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendToTelegram(repos []parse.RepoList) {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error Loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	chatIDstr := os.Getenv("CHAT_ID")

	chatID, chaterr := strconv.ParseInt(chatIDstr, 10, 64)
	if chaterr != nil {
		log.Fatalf("Invalid CHAT_ID: %v", chaterr)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	for _, repo := range repos {
		msg_formatted := fmt.Sprintf("New potential leak on Github\n%s", repo.URL)
		msg := tgbotapi.NewMessage(chatID, msg_formatted)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Erreur envoi Telegram pour %s: %v\n", repo.Name, err)
		}
	}
}
