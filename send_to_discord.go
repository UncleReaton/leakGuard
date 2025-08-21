package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"leakGuard/parse"

	"github.com/joho/godotenv"
)

func sendToDiscord(repos []parse.RepoList) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting home directory: %sv\n", err)
		return
	}

	configDir := filepath.Join(homedir, ".config", "leakGuard")
	filePath := filepath.Join(configDir, ".env")

	enverr := godotenv.Load(filePath)
	if enverr != nil {
		log.Fatal("Error Loading .env file")
	}

	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK")

	if discordWebhookURL == "" {
		log.Fatal("DISCORD_WEBHOOK not set in .env")
	}

	for _, repo := range repos {
		msg_formatted := fmt.Sprintf("New potential leak on Github\n%s", repo.URL)
		if err := postToDiscord(discordWebhookURL, msg_formatted); err != nil {
			log.Printf("Erreur envoi Discord pour %s: %v\n", repo.Name, err)
		}
	}
}

func postToDiscord(webhookURL, message string) error {
	payload := map[string]string{
		"content": message,
	}
	body, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("discord webhook returned status %d", resp.StatusCode)
	}

	return nil
}
