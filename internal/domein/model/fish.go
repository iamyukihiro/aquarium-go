package model

import (
	enum2 "aquarium/internal/domein/model/enum"
)

type FishImpl struct {
	Name               string                   `json:"name"`
	Breed              BreedImpl                `json:"breed"`
	ConditionLevelType enum2.ConditionLevelType `json:"conditionLevelType"`
	HungerLevelType    enum2.HungerLevelType    `json:"hungerNameType"`
}

type BreedImpl struct {
	FishType      enum2.FishType      `json:"fishType"`
	BreedNameType enum2.BreedNameType `json:"breedNameType"`
}

func CreateFish(name string, fishType enum2.FishType, breedNameType enum2.BreedNameType) FishImpl {
	return FishImpl{
		Name:               name,
		Breed:              BreedImpl{FishType: fishType, BreedNameType: breedNameType},
		ConditionLevelType: enum2.Fine,
		HungerLevelType:    enum2.Stuffed,
	}
}
