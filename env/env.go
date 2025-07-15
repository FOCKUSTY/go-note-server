package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func Get(name string) string {
	Init()

	value, exists := os.LookupEnv(name)

	if !exists {
		return ""
	}

	return value
}
