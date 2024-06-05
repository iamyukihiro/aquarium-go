package logic

import (
	"math/rand"
	"time"
)

type ProbabilityImpl struct{}

func NewProbability() *ProbabilityImpl {
	return &ProbabilityImpl{}
}

func (p *ProbabilityImpl) Calc(percentage int) bool {
	rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	return randomNumber <= percentage
}
