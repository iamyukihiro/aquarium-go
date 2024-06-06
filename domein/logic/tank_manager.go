package logic

import (
	"aquarium/domein/model"
	"aquarium/domein/model/enum"
	"encoding/json"
)

type TankManager interface {
	InjectTankManager() Filer
	Init()
}

type TankManagerImpl struct {
	TankManager
	filer Filer
}

func InjectFiler() Filer {
	return NewFiler()
}

func NewTankManager() *TankManagerImpl {
	return &TankManagerImpl{filer: InjectFiler()}
}

func (tm *TankManagerImpl) Init() {
	medaka := model.CreateFish(NewNicknameGenerator().Generate(), enum.MEDAKA, enum.HI_MEDAKA)
	bass := model.CreateFish(NewNicknameGenerator().Generate(), enum.BASS, enum.LARGE_MOUSE)

	t := model.TankImpl{
		Name:     "テスト水槽",
		FishList: []model.FishImpl{medaka, bass},
	}

	file, _ := json.MarshalIndent(t, "", " ")
	_ = tm.filer.WriteFile(file)
}

func (tm *TankManagerImpl) Load() model.TankImpl {
	jsonRaw, _ := tm.filer.ReadFile()

	var t model.TankImpl
	_ = json.Unmarshal(jsonRaw, &t)

	return t
}

func (tm *TankManagerImpl) Save(tank model.TankImpl) {
	file, _ := json.MarshalIndent(tank, "", " ")
	_ = tm.filer.WriteFile(file)
}
