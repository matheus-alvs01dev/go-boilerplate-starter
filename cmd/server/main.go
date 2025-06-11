package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	fmt.Println("Configuration loaded successfully!")
	fmt.Println("Server: ", config.GetServerConfig())
	fmt.Println("Environment: ", config.GetEnv())
}
