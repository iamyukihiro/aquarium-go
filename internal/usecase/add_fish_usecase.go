package usecase

import (
	"aquarium/internal/domein/logic"
	"aquarium/internal/domein/model"
	"aquarium/internal/domein/model/enum"
)

type AddFishUseCase struct {
	tankManager logic.TankManagerInterface
}

func NewAddFishUseCase() *AddFishUseCase {
	return &AddFishUseCase{tankManager: InjectTankManager()}
}

func (afu *AddFishUseCase) AddFish() {
	tank := afu.tankManager.Load()
	tank.AddFish(gachaFish())
	afu.tankManager.Save(tank)
}

func gachaFish() model.Fish {
	if logic.NewProbability().Calc(3) {
		return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.Bass, enum.LargeMouse)
	}

	return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.Medaka, enum.HiMedaka)
}
