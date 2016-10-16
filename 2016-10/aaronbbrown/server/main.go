package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	zmq "github.com/pebbe/zmq4"
)

type Score struct {
	Player [2]int
	Ties   int
}

// get an environment variable as an integer with a default
func getEnvNDefault(key string, defValue int) (n int, err error) {
	s, exists := os.LookupEnv(key)
	if exists {
		n, err = strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("%s must be an integer.  Got %s", key, s)
		}
	} else {
		n = 10
	}
	return n, nil
}

func main() {
	//	score := Score{}

	gameCount, err := getEnvNDefault("GAMES", 10)
	if err != nil {
		log.Fatal(err)
	}
	port, err := getEnvNDefault("PORT", 5555)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Games to play: %d\n", gameCount)
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.PAIR)
	defer context.Term()
	defer socket.Close()

	bindStr := fmt.Sprintf("tcp://*:%d", port)
	socket.Bind(bindStr)

	// Wait for messages
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		// send reply back to client
		reply := fmt.Sprintf("World")
		socket.Send(reply, 0)
	}
}
