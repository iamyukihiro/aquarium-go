package enum

type ConditionLevelType string

const (
	Fine = ConditionLevelType("Fine")
	Bad  = ConditionLevelType("Bad")
	Dead = ConditionLevelType("Dead")
)
