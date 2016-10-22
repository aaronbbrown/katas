package main

import (
	"fmt"
	"log"

	"github.com/aaronbbrown/rps/rps"
	zmq "github.com/pebbe/zmq4"
)

func ZmqServer(games int, port int, control chan int) {
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
			strategy := &rps.RandomStrategy{}
			game := NewZmqGame(socket, i, strategy)
			outcome, err := game.Play(rps.You)
			if err != nil {
				log.Println(err)
				break
			}

			outcome, err = game.Outcome()
			if err != nil {
				fmt.Println(err)
				break
			}
			outcome.UpdateScore(&score)

			fmt.Print(game.String())
			fmt.Printf("Winner:\t%s\n", outcome.String())
			fmt.Printf("Score:\t%s\n\n", score.String())
		}
		// need this extra receive because the client is optimistic and doesn't
		// know how many games we're going to play.
		socket.Recv(0)
		socket.Send("end", 0)

		fmt.Printf("Overall Winner: %s\n\n", score.Winner().String())
	}
	control <- 1
}
