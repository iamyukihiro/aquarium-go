package logic

import (
	model2 "aquarium/internal/domein/model"
	enum2 "aquarium/internal/domein/model/enum"
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

func getTestTank() model2.TankImpl {
	medaka := model2.CreateFish("medaka1", enum2.Medaka, enum2.HiMedaka)
	bass := model2.CreateFish("bass1", enum2.Bass, enum2.LargeMouse)
	t := model2.TankImpl{
		Name:     "テスト水槽",
		FishList: []model2.FishImpl{medaka, bass},
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
