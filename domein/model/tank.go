package model

type TankImpl struct {
	Name     string     `json:"name"`
	FishList []FishImpl `json:"fishList"`
}

func (t *TankImpl) AddFish(fish FishImpl) {
	newFishList := append(t.FishList, fish)
	t.FishList = newFishList
}
