package logic

import "aquarium/internal/domein/model/enum"

type GenderGenerator struct {
	probability ProbabilityInterface
	genderBias  int
}

func NewGenderGenerator() *GenderGenerator {
	return &GenderGenerator{
		probability: NewProbability(),
		genderBias:  50,
	}
}

func (gg *GenderGenerator) Generate() enum.GenderType {
	if gg.probability.Calc(gg.genderBias) {
		return enum.Male
	}

	return enum.Female
}
