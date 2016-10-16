package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aaronbbrown/rps/rps"

	zmq "github.com/pebbe/zmq4"
)

func Client() {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()
	address := os.Getenv("ADDRESS")
	socket.Connect(address)

	score := rps.Score{}
	i := 0
	for {
		i++
		game := &rps.Game{}

		strategy := &rps.RandomStrategy{}
		me := strategy.Throw()
		socket.Send(me.String(), 0)
		game.Throw(rps.Me, me)

		reply, _ := socket.Recv(0)
		if reply == "end" {
			break
		}

		you, err := rps.ThrowTypeFromString(reply)
		if err != nil {
			log.Fatalf("%v", err)
		}
		game.Throw(rps.You, you)

		outcome, err := game.Outcome()
		if err != nil {
			log.Fatalf("%v", err)
		}

		outcome.UpdateScore(&score)

		fmt.Printf("Game:\t%d\n", i)
		fmt.Printf("Me:\t%s\n", me)
		fmt.Printf("You:\t%s\n", you)
		fmt.Printf("Winner:\t%s\n", outcome.String())
		fmt.Printf("Score:\t%s\n\n", score.String())
	}
}
