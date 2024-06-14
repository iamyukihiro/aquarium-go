package model

import (
	"aquarium/internal/domein/model/enum"
)

type Fish struct {
	Name               string                  `json:"name"`
	Breed              Breed                   `json:"breed"`
	ConditionLevelType enum.ConditionLevelType `json:"conditionLevelType"`
	HungerLevelType    enum.HungerLevelType    `json:"hungerNameType"`
}

type Breed struct {
	FishType      enum.FishType      `json:"fishType"`
	BreedNameType enum.BreedNameType `json:"breedNameType"`
}

func CreateFish(name string, fishType enum.FishType, breedNameType enum.BreedNameType) Fish {
	return Fish{
		Name:               name,
		Breed:              Breed{FishType: fishType, BreedNameType: breedNameType},
		ConditionLevelType: enum.Fine,
		HungerLevelType:    enum.Stuffed,
	}
}
