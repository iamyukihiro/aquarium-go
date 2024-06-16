package logic

import (
	"aquarium/internal/domein/model"
	"aquarium/internal/domein/model/enum"
	"encoding/json"
)

type TankManagerInterface interface {
	Init()
	Load() model.Tank
	Save(tank model.Tank)
}

type TankManager struct {
	TankManagerInterface
	filer FilerInterface
}

func NewTankManager() *TankManager {
	return &TankManager{filer: NewFiler()}
}

func (tm *TankManager) Init() {
	medaka := model.CreateFish(NewNicknameGenerator().Generate(), enum.Medaka, enum.HiMedaka, NewGenderGenerator().Generate())
	bass := model.CreateFish(NewNicknameGenerator().Generate(), enum.Bass, enum.LargeMouse, NewGenderGenerator().Generate())

	t := model.Tank{
		Name:     "テスト水槽",
		FishList: []model.Fish{medaka, bass},
	}

	file, _ := json.MarshalIndent(t, "", " ")
	_ = tm.filer.WriteFile(file)
}

func (tm *TankManager) Load() model.Tank {
	jsonRaw, _ := tm.filer.ReadFile()

	var t model.Tank
	_ = json.Unmarshal(jsonRaw, &t)

	return t
}

func (tm *TankManager) Save(tank model.Tank) {
	file, _ := json.MarshalIndent(tank, "", " ")
	_ = tm.filer.WriteFile(file)
}
