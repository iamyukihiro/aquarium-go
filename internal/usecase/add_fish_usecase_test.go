package usecase

import (
	"aquarium/internal/domein/logic"
	"aquarium/internal/domein/logic/mock"
	"aquarium/internal/domein/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTankManagerAddFishUseCase struct {
	mocklogic.MockTankManager
	loadCalled bool
	saveCalled bool
	resultTank model.Tank
}

func (m *MockTankManagerAddFishUseCase) Load() model.Tank {
	m.loadCalled = true
	t := model.Tank{
		Name:     "テスト水槽",
		FishList: []model.Fish{},
	}
	return t
}

func (m *MockTankManagerAddFishUseCase) Save(tank model.Tank) {
	m.resultTank = tank
	m.saveCalled = true
}

func (m *MockTankManagerAddFishUseCase) InjectFiler() logic.Filer {
	panic("呼ばれません")
}

func (m *MockTankManagerAddFishUseCase) Init() {
	m.loadCalled = true
}

func TestAddFishUseCase_InitTank(t *testing.T) {
	mtm := &MockTankManagerAddFishUseCase{}
	SUT := &AddFishUseCase{tankManager: mtm}
	SUT.AddFish()

	assert.True(t, mtm.loadCalled, "MockTankManager の Load() が呼ばれていません")
	assert.True(t, mtm.saveCalled, "MockTankManager の Save() が呼ばれていません")
	assert.Equal(t, len(mtm.resultTank.FishList), 1)
}