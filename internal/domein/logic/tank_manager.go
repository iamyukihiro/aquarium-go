package logic

import (
	model2 "aquarium/internal/domein/model"
	enum2 "aquarium/internal/domein/model/enum"
	"encoding/json"
)

type TankManager interface {
	Init()
	Load() model2.TankImpl
	Save(tank model2.TankImpl)
}

type TankManagerImpl struct {
	TankManager
	filer Filer
}

func NewTankManager() *TankManagerImpl {
	return &TankManagerImpl{filer: InjectFiler()}
}

func (tm *TankManagerImpl) Init() {
	medaka := model2.CreateFish(NewNicknameGenerator().Generate(), enum2.Medaka, enum2.HiMedaka)
	bass := model2.CreateFish(NewNicknameGenerator().Generate(), enum2.Bass, enum2.LargeMouse)

	t := model2.TankImpl{
		Name:     "テスト水槽",
		FishList: []model2.FishImpl{medaka, bass},
	}

	file, _ := json.MarshalIndent(t, "", " ")
	_ = tm.filer.WriteFile(file)
}

func (tm *TankManagerImpl) Load() model2.TankImpl {
	jsonRaw, _ := tm.filer.ReadFile()

	var t model2.TankImpl
	_ = json.Unmarshal(jsonRaw, &t)

	return t
}

func (tm *TankManagerImpl) Save(tank model2.TankImpl) {
	file, _ := json.MarshalIndent(tank, "", " ")
	_ = tm.filer.WriteFile(file)
}
