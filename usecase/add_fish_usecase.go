package usecase

import (
	"aquarium/domein/model"
	"aquarium/domein/model/enum"
)

func AddFish() {
	tank := model.NewTankManager().Load()
	bass := model.CreateFish("ブラックバス", enum.BASS, enum.LARGE_MOUSE)

	tank.AddFish(bass)
	model.NewTankManager().Save(tank)
}
