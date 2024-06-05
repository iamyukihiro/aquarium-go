package model

import (
	"aquarium/domein/model/enum"
	"encoding/json"
	"os"
)

type TankManagerImpl struct {
}

func NewTankManager() *TankManagerImpl {
	return &TankManagerImpl{}
}

func (r *TankManagerImpl) Init() {
	medaka := CreateFish("メダカ", enum.MEDAKA, enum.HI_MEDAKA)
	bass := CreateFish("ブラックバス", enum.BASS, enum.LARGE_MOUSE)

	t := Tank{
		Name:     "テスト",
		FishList: []FishImpl{medaka, bass},
	}

	file, _ := json.MarshalIndent(t, "", " ")
	_ = os.WriteFile("save/tank.json", file, 0644)
}

func (r *TankManagerImpl) Load() Tank {
	jsonRaw, _ := os.ReadFile("save/tank.json")

	var t Tank
	_ = json.Unmarshal(jsonRaw, &t)

	return t
}

func (r *TankManagerImpl) Save(tank Tank) {
	file, _ := json.MarshalIndent(tank, "", " ")
	_ = os.WriteFile("save/tank.json", file, 0644)
}
