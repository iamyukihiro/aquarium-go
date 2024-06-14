package usecase

import (
	"aquarium/internal/domein/logic"
)

type InitTankUseCase struct {
	tankManager logic.TankManagerInterface
}

func NewInitTankUseCase() *InitTankUseCase {
	return &InitTankUseCase{tankManager: InjectTankManager()}
}

func (itu *InitTankUseCase) InitTank() {
	itu.tankManager.Init()
}
