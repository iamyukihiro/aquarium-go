package logic

import (
	"math/rand"
	"time"
)

type Probability struct{}

func NewProbability() *Probability {
	return &Probability{}
}

func (p *Probability) Calc(percentage int) bool {
	rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	return randomNumber <= percentage
}
