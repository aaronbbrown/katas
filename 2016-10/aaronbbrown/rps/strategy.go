package rps

import (
	"math/rand"
	"time"
)

type Strategy interface {
	Throw() ThrowType
}

type RandomStrategy struct{}

func (r RandomStrategy) Throw() ThrowType {
	rand.Seed(time.Now().UnixNano())
	return ThrowType(rand.Intn(3))
}
