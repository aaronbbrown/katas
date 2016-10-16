package main

import "os"

func main() {
	mode := os.Getenv("MODE")

	switch mode {
	case "client":
		Client()
	default:
		Server()
	}
}
