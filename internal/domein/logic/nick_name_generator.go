package logic

import (
	"math/rand"
	"time"
)

type NicknameGenerator struct {
	fishNicknames []string
}

func NewNicknameGenerator() *NicknameGenerator {
	return &NicknameGenerator{
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

func (ng *NicknameGenerator) Generate() string {
	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(ng.fishNicknames))

	return ng.fishNicknames[randomIndex]
}
