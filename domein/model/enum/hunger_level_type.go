package enum

type HungerLevelType string

const (
	Starving     = HungerLevelType("STARVING")
	VeryHungry   = HungerLevelType("VERY_HUNGRY")
	Hungry       = HungerLevelType("HUNGRY")
	LittleHunger = HungerLevelType("LITTLE_HUNGER")
	Stuffed      = HungerLevelType("STUFFED")
)
