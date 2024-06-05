package usecase

import (
	"aquarium/domein/logic"
	"aquarium/domein/model"
	"aquarium/domein/model/enum"
)

func AddFish() {
	tank := model.NewTankManager().Load()

	tank.AddFish(gachaFish())
	model.NewTankManager().Save(tank)
}

func gachaFish() model.FishImpl {
	if logic.NewProbability().Calc(3) {
		return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.BASS, enum.LARGE_MOUSE)
	}

	return model.CreateFish(logic.NewNicknameGenerator().Generate(), enum.MEDAKA, enum.HI_MEDAKA)
}
