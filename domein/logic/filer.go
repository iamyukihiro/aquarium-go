package logic

import (
	"fmt"
	"os"
)

const TankDir = "save/tank.json"

type Filer interface {
	WriteFile(file []byte) error
	ReadFile() ([]byte, error)
}

type FilerImpl struct {
	filer Filer
}

func NewFiler() *FilerImpl {
	return &FilerImpl{}
}

func (f *FilerImpl) WriteFile(file []byte) error {
	err := os.WriteFile(TankDir, file, 0644)
	if err != nil {
		return fmt.Errorf("ファイル書き込み時にエラーが発生しました: %v", err)
	}

	return nil
}

func (f *FilerImpl) ReadFile() ([]byte, error) {
	file, err := os.ReadFile(TankDir)
	if err != nil {
		return nil, fmt.Errorf("ファイル読み込み時にエラーが発生しました: %v", err)
	}

	return file, nil
}
