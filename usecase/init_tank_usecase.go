package usecase

import (
	"aquarium/domein/model"
)

func InitTank() {
	model.NewTankManager().Init()
}
