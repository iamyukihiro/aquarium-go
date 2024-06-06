package usecase

import (
	"aquarium/domein/logic"
	"aquarium/domein/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTankManager struct {
	initCalled bool
	InitTankUseCaseImpl
}

func (m *mockTankManager) Load() model.TankImpl {
	panic("呼ばれません")
}

func (m *mockTankManager) Save(tank model.TankImpl) {
	panic("呼ばれません")
}

func (m *mockTankManager) InjectTankManager() logic.Filer {
	panic("呼ばれません")
}

func (m *mockTankManager) Init() {
	m.initCalled = true
}

func TestInitTankUseCase_InitTank(t *testing.T) {
	mtm := &mockTankManager{}
	SUT := &InitTankUseCaseImpl{tankManager: mtm}
	SUT.InitTank()

	assert.True(t, mtm.initCalled, "MockTankManager の Init() が呼ばれていません")
}
