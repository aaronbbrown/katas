package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"../rps"

)

type Score struct {
	Player [2]int
	Ties   int
}

func main() {
	var gameCount int
	var err error
	score := Score{}

	gameCountStr, exists := os.LookupEnv("GAMES")
	if exists {
		gameCount, err = strconv.Atoi(gameCountStr)
		if err != nil {
			log.Fatalf("GAMES must be an integer.  Got %v", gameCountStr)
		}
	} else {
		gameCount = 10
	}

	fmt.Printf("Games to play: %d\n", gameCount)

	for i := 0; i < gameCount; i++ {
		game := &rps.Game{}
		strategy := &rps.RandomStrategy{}

		p1 := strategy.Throw()
		p2 := strategy.Throw()

		game.Throw(rps.Player1, p1)
		fmt.Printf("Player 1 threw %s\n", p1)

		game.Throw(rps.Player2, p2)
		fmt.Printf("Player 2 threw %s\n", p2)

		outcome, err := game.Outcome()
		if err != nil {
			log.Fatalf("Error with game: %v", err)
		}

		if outcome.Tie {
			fmt.Printf("Tie!\n")
			score.Ties++
		} else {
			fmt.Printf("The winner is player %d\n", outcome.Winner+1)
			score.Player[outcome.Winner]++
		}
	}

	fmt.Printf("Final score: %v\n", score)
}
