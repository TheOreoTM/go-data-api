package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/theoreotm/go-data-api/application"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Get environment variables
	redisPortStr := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisHost := os.Getenv("REDIS_HOST")

	// Convert the port to an integer
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		fmt.Println("Error converting REDIS_PORT to integer:", err)
		return
	}

	app := application.New(redisPassword, redisHost, redisPort, 0)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
