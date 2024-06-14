package logic

import (
	"aquarium/internal/domein/model"
	"aquarium/internal/domein/model/enum"
	"encoding/json"
)

type TankManager interface {
	Init()
	Load() model.TankImpl
	Save(tank model.TankImpl)
}

type TankManagerImpl struct {
	TankManager
	filer Filer
}

func NewTankManager() *TankManagerImpl {
	return &TankManagerImpl{filer: InjectFiler()}
}

func (tm *TankManagerImpl) Init() {
	medaka := model.CreateFish(NewNicknameGenerator().Generate(), enum.Medaka, enum.HiMedaka)
	bass := model.CreateFish(NewNicknameGenerator().Generate(), enum.Bass, enum.LargeMouse)

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
