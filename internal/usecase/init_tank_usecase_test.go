package usecase

import (
	"aquarium/internal/domein/logic"
	"aquarium/internal/domein/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTankManager struct {
	initCalled bool
	logic.TankManagerInterface
}

func (m *mockTankManager) Load() model.Tank {
	panic("呼ばれません")
}

func (m *mockTankManager) Save(tank model.Tank) {
	panic("呼ばれません")
}

func (m *mockTankManager) InjectFiler() logic.Filer {
	panic("呼ばれません")
}

func (m *mockTankManager) Init() {
	m.initCalled = true
}

func TestInitTankUseCase_InitTank(t *testing.T) {
	mtm := &mockTankManager{}
	SUT := &InitTankUseCase{tankManager: mtm}
	SUT.InitTank()

	assert.True(t, mtm.initCalled, "MockTankManager の Init() が呼ばれていません")
}
