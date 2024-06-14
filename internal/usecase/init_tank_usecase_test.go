package usecase

import (
	"aquarium/internal/domein/logic"
	"aquarium/internal/domein/logic/mock"
	"aquarium/internal/domein/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTankManagerInitTankUseCase struct {
	mocklogic.MockTankManager
	initCalled bool
}

func (m *MockTankManagerInitTankUseCase) Load() model.Tank {
	panic("呼ばれません")
}

func (m *MockTankManagerInitTankUseCase) Save(tank model.Tank) {
	panic("呼ばれません")
}

func (m *MockTankManagerInitTankUseCase) InjectFiler() logic.Filer {
	panic("呼ばれません")
}

func (m *MockTankManagerInitTankUseCase) Init() {
	m.initCalled = true
}

func TestInitTankUseCase_InitTank(t *testing.T) {
	mtm := &MockTankManagerInitTankUseCase{}
	SUT := &InitTankUseCase{tankManager: mtm}
	SUT.InitTank()

	assert.True(t, mtm.initCalled, "MockTankManager の Init() が呼ばれていません")
}
