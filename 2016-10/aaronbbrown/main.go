package main

import (
	"log"
	"os"
)

func main() {
	mode := os.Getenv("MODE")

	switch mode {
	case "client":
		address := os.Getenv("ADDRESS")
		ZmqClient(address)

	default:
		control := make(chan int)
		games, err := GetEnvNDefault("GAMES", 10)
		if err != nil {
			log.Fatal(err)
		}
		port, err := GetEnvNDefault("PORT", 5555)
		if err != nil {
			log.Fatal(err)
		}

		ZmqServer(games, port, control)
	}
}
