package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"

	zmq "github.com/pebbe/zmq4"
)

func ZmqClient(address string) {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()
	socket.Connect(address)

	score := rps.Score{}
	i := 0
	for {
		i++
		strategy := &rps.RandomStrategy{}
		game := NewZmqGame(socket, i, strategy)
		outcome, err := game.Play(rps.Me)
		if err != nil {
			log.Fatal(err)
		}
		if outcome.End {
			break
		}

		outcome.UpdateScore(&score)

		fmt.Println(game.String())
		fmt.Printf("Winner:\t%s\n", outcome.String())
		fmt.Printf("Score:\t%s\n", score.String())
	}
	fmt.Printf("\nOverall Winner: %s\n\n", score.Winner().String())
}
