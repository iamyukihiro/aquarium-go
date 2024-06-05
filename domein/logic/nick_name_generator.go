package logic

import (
	"math/rand"
	"time"
)

type NicknameGeneratorImpl struct {
	fishNicknames []string
}

func NewNicknameGenerator() *NicknameGeneratorImpl {
	return &NicknameGeneratorImpl{
		fishNicknames: []string{
			"ジョン",
			"たけし",
			"マイク",
			"マイケル",
			"ジョニー",
			"ボブ",
			"ありさわ",
			"ゆきひろ",
		},
	}
}

func (ng *NicknameGeneratorImpl) Generate() string {
	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(ng.fishNicknames))

	return ng.fishNicknames[randomIndex]
}
