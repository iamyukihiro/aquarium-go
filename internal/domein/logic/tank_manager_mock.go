package logic

type MockTankManager struct {
	TankManagerInterface
	LoadCalled bool
	SaveCalled bool
	InitCalled bool
}
