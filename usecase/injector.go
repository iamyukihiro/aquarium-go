package usecase

import (
	"aquarium/domein/logic"
)

func InjectTankManager() logic.TankManager {
	return logic.NewTankManager()
}
