package logic

import (
	"aquarium/internal/domein/model"
	"aquarium/internal/domein/model/enum"
	"encoding/json"
	"reflect"
	"testing"
)

type MockFilerImpl struct {
	MockFiler
}

func (f *MockFilerImpl) WriteFile(file []byte) error {
	return nil
}

func (f *MockFilerImpl) ReadFile() ([]byte, error) {
	t := getTestTank()

	file, _ := json.MarshalIndent(t, "", " ")

	return file, nil
}

func getTestTank() model.Tank {
	medaka := model.CreateFish("medaka1", enum.Medaka, enum.HiMedaka)
	bass := model.CreateFish("bass1", enum.Bass, enum.LargeMouse)
	t := model.Tank{
		Name:     "テスト水槽",
		FishList: []model.Fish{medaka, bass},
	}

	return t
}

func TestTankManager_Init(t *testing.T) {
	t.Skip()
}

func TestTankManager_Load(t *testing.T) {
	expected := getTestTank()

	mockFiler := &MockFilerImpl{}
	tm := &TankManager{filer: mockFiler}
	actual := tm.Load()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Init メソッドでのデータ書き込みが正しくありません。期待されるデータ: %+v, 実際のデータ: %+v", expected, actual)
	}
}

func TestTankManager_Save(t *testing.T) {
	t.Skip()
}
