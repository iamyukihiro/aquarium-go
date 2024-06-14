package logic

import (
	"math/rand"
	"testing"
	"time"
)

func TestNicknameGeneratorImpl_Generate(t *testing.T) {
	SUT := NewNicknameGenerator()

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		nickname := SUT.Generate()
		if !contains(SUT.fishNicknames, nickname) {
			t.Errorf("ありえないニックネームです。 : %s", nickname)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
