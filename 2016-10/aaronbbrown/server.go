package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"
	zmq "github.com/pebbe/zmq4"
)

func Server() {
	gameCount, err := GetEnvNDefault("GAMES", 10)
	if err != nil {
		log.Fatal(err)
	}
	port, err := GetEnvNDefault("PORT", 5555)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("Games to play: %d\n", gameCount)
		socket, _ := zmq.NewSocket(zmq.PAIR)
		defer socket.Close()

		bindStr := fmt.Sprintf("tcp://*:%d", port)
		socket.Bind(bindStr)
		fmt.Printf("Socket: %s\n\n", bindStr)

		score := rps.Score{}

		// Wait for messages
		for i := 1; i <= gameCount; i++ {
			// GAME. ON!
			msg, _ := socket.Recv(0)

			game := &rps.Game{}

			strategy := &rps.RandomStrategy{}
			me := strategy.Throw()
			socket.Send(me.String(), 0)
			game.Throw(rps.Me, me)

			you, err := rps.ThrowTypeFromString(msg)
			if err != nil {
				break
			}
			game.Throw(rps.You, you)

			outcome, err := game.Outcome()
			if err != nil {
				break
			}
			outcome.UpdateScore(&score)

			fmt.Printf("Game:\t%d\n", i)
			fmt.Printf("Me:\t%s\n", me)
			fmt.Printf("You:\t%s\n", you)
			fmt.Printf("Winner:\t%s\n", outcome.String())
			fmt.Printf("Score:\t%s\n\n", score.String())
		}
		socket.Send("end", 0)
		socket.Close()
	}
}
