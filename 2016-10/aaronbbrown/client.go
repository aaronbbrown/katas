package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"

	zmq "github.com/pebbe/zmq4"
)

func Client(address string) {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()
	socket.Connect(address)

	score := rps.Score{}
	i := 0
	for {
		i++
		game := &rps.Game{}

		strategy := &rps.RandomStrategy{}
		me := strategy.Throw()
		socket.Send(me.String(), 0)
		fmt.Printf("Sent %v\n", me.String())
		game.Throw(rps.Me, me)

		reply, err := socket.Recv(0)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Printf("Received: %v\n", reply)

		if reply == "end" {
			break
		}

		you, err := rps.ThrowTypeFromString(reply)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		game.Throw(rps.You, you)

		outcome, err := game.Outcome()
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		outcome.UpdateScore(&score)

		fmt.Printf("Game:\t%d\n", i)
		fmt.Printf("Me:\t%s\n", me)
		fmt.Printf("You:\t%s\n", you)
		fmt.Printf("Winner:\t%s\n", outcome.String())
		fmt.Printf("Score:\t%s\n\n", score.String())
	}
}
