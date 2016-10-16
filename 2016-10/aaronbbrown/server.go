package main

import (
	"fmt"

	"github.com/aaronbbrown/rps/rps"
	zmq "github.com/pebbe/zmq4"
)

func Server(games int, port int, control chan int) {
	socket, _ := zmq.NewSocket(zmq.PAIR)
	defer socket.Close()

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	socket.Bind(bindStr)
	fmt.Printf("Socket: %s\n\n", bindStr)

	for {
		fmt.Printf("Games to play: %d\n", games)
		score := rps.Score{}

		// Wait for messages
		for i := 1; i <= games; i++ {
			// GAME. ON!
			msg, err := socket.Recv(0)
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			fmt.Printf("Received: %v", msg)

			game := &rps.Game{}

			strategy := &rps.RandomStrategy{}
			me := strategy.Throw()
			socket.Send(me.String(), 0)
			fmt.Printf("Sent %v\n", me.String())
			game.Throw(rps.Me, me)

			you, err := rps.ThrowTypeFromString(msg)
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			game.Throw(rps.You, you)

			outcome, err := game.Outcome()
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			outcome.UpdateScore(&score)

			fmt.Printf("Game:\t%d\n", i)
			fmt.Printf("Me:\t%s\n", me)
			fmt.Printf("You:\t%s\n", you)
			fmt.Printf("Winner:\t%s\n", outcome.String())
			fmt.Printf("Score:\t%s\n\n", score.String())
		}
		// need this extra receive because the client is optimistic and doesn't
		// know how many games we're going to play.
		socket.Recv(0)
		socket.Send("end", 0)
	}
	control <- 1
}
