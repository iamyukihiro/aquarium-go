package model

type Tank struct {
	Name     string `json:"name"`
	FishList []Fish `json:"fishList"`
}

func (t *Tank) AddFish(fish Fish) {
	newFishList := append(t.FishList, fish)
	t.FishList = newFishList
}
