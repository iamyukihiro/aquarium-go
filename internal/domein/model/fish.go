package model

import (
	"aquarium/internal/domein/model/enum"
)

type FishImpl struct {
	Name               string                  `json:"name"`
	Breed              BreedImpl               `json:"breed"`
	ConditionLevelType enum.ConditionLevelType `json:"conditionLevelType"`
	HungerLevelType    enum.HungerLevelType    `json:"hungerNameType"`
}

type BreedImpl struct {
	FishType      enum.FishType      `json:"fishType"`
	BreedNameType enum.BreedNameType `json:"breedNameType"`
}

func CreateFish(name string, fishType enum.FishType, breedNameType enum.BreedNameType) FishImpl {
	return FishImpl{
		Name:               name,
		Breed:              BreedImpl{FishType: fishType, BreedNameType: breedNameType},
		ConditionLevelType: enum.Fine,
		HungerLevelType:    enum.Stuffed,
	}
}
