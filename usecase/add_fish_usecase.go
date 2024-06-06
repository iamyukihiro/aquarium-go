package usecase

import (
	"aquarium/domein/logic"
	"aquarium/domein/model"
	"aquarium/domein/model/enum"
)

type AddFishUseCaseImpl struct {
	tankManager logic.TankManager
}

func NewAddFishUseCase() *AddFishUseCaseImpl {
	return &AddFishUseCaseImpl{tankManager: InjectTankManager()}
}

func (afu *AddFishUseCaseImpl) AddFish() {
	tank := afu.tankManager.Load()
	tank.AddFish(gachaFish())
	afu.tankManager.Save(tank)
}

func gachaFish() model.FishImpl {
	if logic.NewProbability().Calc(3) {
		return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.BASS, enum.LARGE_MOUSE)
	}

	return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.MEDAKA, enum.HI_MEDAKA)
}
