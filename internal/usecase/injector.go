package usecase

import (
	"aquarium/internal/domein/logic"
)

func InjectTankManager() logic.TankManagerInterface {
	return logic.NewTankManager()
}
