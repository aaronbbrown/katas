package rps

import (
	"fmt"
)

const (
	Player1 int = iota
	Player2
)

type GameOutcome struct {
	Tie    bool
	Winner int
}

type Game struct {
	Throws [2]Throw
}

func (g *Game) Outcome() (*GameOutcome, error) {
	outcome := &GameOutcome{}
	if !g.Throws[Player1].Thrown || !g.Throws[Player2].Thrown {
		return outcome, fmt.Errorf("Both throws haven't been made")
	}

	if g.Throws[Player1].Type == g.Throws[Player2].Type {
		outcome.Tie = true
	} else if saneModInt(int(g.Throws[Player1].Type-1), 3) == int(g.Throws[Player2].Type) {
		outcome.Winner = Player1
	} else {
		outcome.Winner = Player2
	}

	return outcome, nil
}

func (g *Game) Throw(player int, tt ThrowType) error {
	if player > len(g.Throws) {
		return fmt.Errorf("Invalid player: %d", player)
	}
	g.Throws[player].Thrown = true
	g.Throws[player].Type = tt
	return nil
}
