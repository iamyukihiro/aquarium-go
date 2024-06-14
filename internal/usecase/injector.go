package usecase

import (
	"aquarium/internal/domein/logic"
)

func InjectTankManager() logic.TankManager {
	return logic.NewTankManager()
}
