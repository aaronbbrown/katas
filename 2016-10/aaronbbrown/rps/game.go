package rps

import (
	"bytes"
	"fmt"
)

const (
	Me int = iota
	You
)

type Game struct {
	Throws [2]Throw
}

func (g *Game) Outcome() (*GameOutcome, error) {
	outcome := &GameOutcome{}
	if !g.Throws[Me].Thrown || !g.Throws[You].Thrown {
		return outcome, fmt.Errorf("Both throws haven't been made")
	}

	if g.Throws[Me].Type == g.Throws[You].Type {
		outcome.Tie = true
	} else if saneModInt(int(g.Throws[Me].Type-1), 3) == int(g.Throws[You].Type) {
		outcome.Winner = Me
	} else {
		outcome.Winner = You
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

type GameOutcome struct {
	Tie    bool
	Winner int
}

func (outcome *GameOutcome) String() string {
	buffer := bytes.NewBufferString("")
	if outcome.Tie {
		buffer.WriteString("Tie")
	} else if outcome.Winner == 0 {
		buffer.WriteString("Me")
	} else {
		buffer.WriteString("You")
	}
	return buffer.String()
}

func (outcome *GameOutcome) UpdateScore(score *Score) {
	if outcome.Tie {
		score.Ties++
	} else {
		score.Player[outcome.Winner]++
	}
}
