package storage

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"leakGuard/parse"
)

func SaveToFile(list []parse.RepoList) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting home directory: %sv\n", err)
		return
	}

	configDir := filepath.Join(homedir, ".config", "leakGuard")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Printf("Error creating config directory: %s\n", err)
		return
	}

	filePath := filepath.Join(configDir, "list.json")
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(list)
	if err != nil {
		log.Printf("Error encoding list: %s\n", err)
	}
}

func LoadList() []parse.RepoList {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Error getting home directory: %sv\n", err)
		return nil
	}

	configDir := filepath.Join(homedir, ".config", "leakGuard")
	filePath := filepath.Join(configDir, "list.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
		return nil
	}

	var list []parse.RepoList

	err = json.Unmarshal(data, &list)
	if err != nil {
		log.Printf("Error decoding json: %v\n", err)
		return nil
	}

	return list
}
