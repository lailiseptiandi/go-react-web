package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func findProjectRoot(dir string, marker string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(dir, marker)); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return "", os.ErrNotExist
}

func LoadEnv() {
	currentWorkDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	marker := "go.mod"
	rootPath, err := findProjectRoot(currentWorkDirectory, marker)
	if err != nil {
		log.Fatalf("Project root containing .env not found from current working directory %s: %v", currentWorkDirectory, err)
	}
	err = godotenv.Load(filepath.Join(rootPath, "app", "env", "app.env"))
	if err != nil {
		log.Fatalf("Error loading .env file from %s: %v", filepath.Join(rootPath, ".env"), err)
	}
}
