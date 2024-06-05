package enum

type ConditionLevelType string

const (
	FINE = ConditionLevelType("Fine")
	BAD  = ConditionLevelType("Bad")
	DEAD = ConditionLevelType("Dead")
)
