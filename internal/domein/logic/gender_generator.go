package logic

import "aquarium/internal/domein/model/enum"

type GenderGenerator struct{}

func NewGenderGenerator() *GenderGenerator {
	return &GenderGenerator{}
}

func (gg *GenderGenerator) Generate() enum.GenderType {
	if NewProbability().Calc(5) {
		return enum.Male
	}

	return enum.Female
}
