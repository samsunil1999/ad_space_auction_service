package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// get current file full path from run time
	_, file, _, _ := runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath := filepath.Join(filepath.Dir(file), "../")

	// load .env file
	err := godotenv.Load(ProjectRootPath + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	loadAppConfig()
}
