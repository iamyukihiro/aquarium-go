package usecase

import (
	"aquarium/domein/logic"
)

type InitTankUseCaseImpl struct {
	tankManager logic.TankManager
}

func InjectTankManager() logic.TankManager {
	return logic.NewTankManager()
}

func NewInitTankUseCase() *InitTankUseCaseImpl {
	return &InitTankUseCaseImpl{tankManager: InjectTankManager()}
}

func (itu *InitTankUseCaseImpl) InitTank() {
	itu.tankManager.Init()
}
