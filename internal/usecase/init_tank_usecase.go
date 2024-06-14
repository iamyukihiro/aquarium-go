package usecase

import (
	"aquarium/internal/domein/logic"
)

type InitTankUseCaseImpl struct {
	tankManager logic.TankManager
}

func NewInitTankUseCase() *InitTankUseCaseImpl {
	return &InitTankUseCaseImpl{tankManager: InjectTankManager()}
}

func (itu *InitTankUseCaseImpl) InitTank() {
	itu.tankManager.Init()
}
