package usecase

import (
	"aquarium/domein/logic"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTankManager struct {
	initCalled bool
	InitTankUseCaseImpl
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
