package model

type Tank struct {
	Name     string     `json:"name"`
	FishList []FishImpl `json:"fishList"`
}

func (t *Tank) AddFish(fish FishImpl) {
	newFishList := append(t.FishList, fish)
	t.FishList = newFishList
}
