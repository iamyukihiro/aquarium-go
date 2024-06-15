package usecase

import (
	"aquarium/internal/domein/logic"
	"aquarium/internal/domein/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTankManagerInitTankUseCase struct {
	logic.MockTankManager
}

func (m *MockTankManagerInitTankUseCase) Load() model.Tank {
	panic("呼ばれません")
}

func (m *MockTankManagerInitTankUseCase) Save(tank model.Tank) {
	panic("呼ばれません")
}

func (m *MockTankManagerInitTankUseCase) Init() {
	m.InitCalled = true
}

func TestInitTankUseCase_InitTank(t *testing.T) {
	mtm := &MockTankManagerInitTankUseCase{}
	SUT := &InitTankUseCase{tankManager: mtm}
	SUT.InitTank()

	assert.True(t, mtm.InitCalled, "MockTankManager の Init() が呼ばれていません")
}
