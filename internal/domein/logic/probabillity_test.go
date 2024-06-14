package logic

import (
	"testing"
)

func TestCalc_0パーセントの時は必ずfalseになること(t *testing.T) {
	SUT := NewProbability()

	result := SUT.Calc(0)
	if result != false {
		t.Errorf("0パーセント時は必ずfalseです。")
	}
}

func TestCalc_100パーセントの時は必ずtrueになること(t *testing.T) {
	SUT := NewProbability()

	result := SUT.Calc(100)
	if result != true {
		t.Errorf("100パーセント時は必ずtrueです。")
	}
}

func TestCalc_確率のバラつきが少ないこと(t *testing.T) {
	SUT := NewProbability()

	successCount := 0
	totalRuns := 100000
	percentage := 30

	for i := 0; i < totalRuns; i++ {
		if SUT.Calc(percentage) {
			successCount++
		}
	}

	actualPercentage := float64(successCount) / float64(totalRuns) * 100
	expectedLowerBound := float64(percentage) - 5 // 許容誤差 -5%
	expectedUpperBound := float64(percentage) + 5 // 許容誤差 +5%

	if actualPercentage < expectedLowerBound || actualPercentage > expectedUpperBound {
		t.Errorf("許容されない確率が算出されました。期待される成功率は約%d%%ですが、実際には%.2f%%でした。", percentage, actualPercentage)
	}
}
