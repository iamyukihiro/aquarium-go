package logic

import (
	"aquarium/domein/model"
	"aquarium/domein/model/enum"
	"encoding/json"
	"reflect"
	"testing"
)

type mockFilerImpl struct {
	Filer
}

func (f *mockFilerImpl) WriteFile(file []byte) error {
	return nil
}

func (f *mockFilerImpl) ReadFile() ([]byte, error) {
	t := getTestTank()

	file, _ := json.MarshalIndent(t, "", " ")

	return file, nil
}

func getTestTank() model.TankImpl {
	medaka := model.CreateFish("medaka1", enum.Medaka, enum.HiMedaka)
	bass := model.CreateFish("bass1", enum.Bass, enum.LargeMouse)
	t := model.TankImpl{
		Name:     "テスト水槽",
		FishList: []model.FishImpl{medaka, bass},
	}

	return t
}

func TestTankManagerImpl_Init(t *testing.T) {
	t.Skip()
}

func TestTankManagerImpl_Load(t *testing.T) {
	expected := getTestTank()

	mockFiler := &mockFilerImpl{}
	tm := &TankManagerImpl{filer: mockFiler}
	actual := tm.Load()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Init メソッドでのデータ書き込みが正しくありません。期待されるデータ: %+v, 実際のデータ: %+v", expected, actual)
	}
}

func TestTankManagerImpl_Save(t *testing.T) {
	t.Skip()
}
