package logic

import (
	"aquarium/internal/domein/model/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("ジェンダーバイアスが100だった場合、Male が返されること", func(t *testing.T) {
		SUT := &GenderGenerator{
			probability: NewProbability(),
			genderBias:  100,
		}
		actual := SUT.Generate()

		assert.Equal(t, enum.Male, actual)
	})

	t.Run("ジェンダーバイアスが0だった場合、Female が返されること", func(t *testing.T) {
		SUT := &GenderGenerator{
			probability: NewProbability(),
			genderBias:  0,
		}
		actual := SUT.Generate()

		assert.Equal(t, enum.Female, actual)
	})
}
