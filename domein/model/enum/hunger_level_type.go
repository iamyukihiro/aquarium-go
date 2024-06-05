package enum

type HungerLevelType string

const (
	STARVING      = HungerLevelType("STARVING")
	VERY_HUNGRY   = HungerLevelType("VERY_HUNGRY")
	HUNGRY        = HungerLevelType("HUNGRY")
	LITTLE_HUNGER = HungerLevelType("LITTLE_HUNGER")
	STUFFED       = HungerLevelType("STUFFED")
)
