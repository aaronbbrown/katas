package rps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Outcome(t *testing.T) {
	// Missing throws
	g := &Game{}
	_, err := g.Outcome()
	assert.NotNil(t, err)

	_, err = g.Outcome()
	assert.NotNil(t, err)

	// Tie
	g = &Game{}
	g.Throw(Player1, Rock)
	g.Throw(Player2, Rock)

	o, err := g.Outcome()
	assert.Nil(t, err)
	assert.True(t, o.Tie)

	// Rock beats Scissors
	g = &Game{}
	g.Throw(Player1, Rock)
	g.Throw(Player2, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player1)

	g = &Game{}
	g.Throw(Player2, Rock)
	g.Throw(Player1, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player2)

	// Scissors beats Paper
	g = &Game{}
	g.Throw(Player1, Scissors)
	g.Throw(Player2, Paper)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player1)

	g = &Game{}
	g.Throw(Player1, Paper)
	g.Throw(Player2, Scissors)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player2)

	// Paper beats Rock
	g = &Game{}
	g.Throw(Player1, Paper)
	g.Throw(Player2, Rock)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player1)

	g = &Game{}
	g.Throw(Player1, Rock)
	g.Throw(Player2, Paper)
	o, err = g.Outcome()
	assert.Nil(t, err)
	assert.False(t, o.Tie)
	assert.Equal(t, o.Winner, Player2)

}
