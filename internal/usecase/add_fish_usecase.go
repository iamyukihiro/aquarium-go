package usecase

import (
	logic2 "aquarium/internal/domein/logic"
	"aquarium/internal/domein/model"
	enum2 "aquarium/internal/domein/model/enum"
)

type AddFishUseCaseImpl struct {
	tankManager logic2.TankManager
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
	if logic2.NewProbability().Calc(3) {
		return model.CreateFish(logic2.NewNicknameGenerator().Generate(), enum2.Bass, enum2.LargeMouse)
	}

	return model.CreateFish(logic2.NewNicknameGenerator().Generate(), enum2.Medaka, enum2.HiMedaka)
}
